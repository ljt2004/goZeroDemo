package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"sync"
	"time"

	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// ================ 核心接口 ================

// CodeGenerator 验证码生成器接口
type CodeGenerator interface {
	Generate() string
}

// CodeStore 验证码存储接口
type CodeStore interface {
	Set(key, code string, expireSeconds int) error
	Get(key string) (string, error)
	Del(key string) error
}

// ================ 实现 ================

// NumberCodeGenerator 数字验证码生成器
type NumberCodeGenerator struct {
	length int
}

var mutex sync.Mutex

func NewNumberCodeGenerator(length int) *NumberCodeGenerator {
	return &NumberCodeGenerator{length: length}
}

func (g *NumberCodeGenerator) Generate() string {
	max := 1
	for i := 0; i < g.length; i++ {
		max *= 10
	}
	return fmt.Sprintf("%0*d", g.length, time.Now().UnixNano()%int64(max))
}

// MixedCodeGenerator 字母数字混合验证码生成器
type MixedCodeGenerator struct {
	length  int
	charset string
}

func NewMixedCodeGenerator(length int) *MixedCodeGenerator {
	return &MixedCodeGenerator{
		length:  length,
		charset: "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz23456789",
	}
}

func (g *MixedCodeGenerator) Generate() string {
	mutex.Lock()
	defer mutex.Unlock()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, g.length)
	for i := range result {
		result[i] = g.charset[r.Intn(len(g.charset))]
	}
	return string(result)
}

// RedisCodeStore Redis存储
type RedisCodeStore struct {
	rds    *redis.Redis
	prefix string
}

func NewRedisCodeStore(rds *redis.Redis, prefix string) *RedisCodeStore {
	return &RedisCodeStore{rds: rds, prefix: prefix}
}

func (s *RedisCodeStore) Set(key, code string, expireSeconds int) error {
	return s.rds.Setex(s.prefix+key, code, expireSeconds)
}

func (s *RedisCodeStore) Get(key string) (string, error) {
	return s.rds.Get(s.prefix + key)
}

func (s *RedisCodeStore) Del(key string) error {
	s.rds.Del(s.prefix + key)
	return nil
}

// EmailSender 邮箱发送器
type EmailSender struct {
	username string
	nick     string
	host     string
	port     int
	password string
	tls      bool
}

func NewEmailSender(username, nick, host string, port int, password string, tls bool) *EmailSender {
	return &EmailSender{
		username: username,
		nick:     nick,
		host:     host,
		port:     port,
		password: password,
		tls:      tls,
	}
}

func (s *EmailSender) Send(to, subject, content string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.nick, s.username)
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(content)

	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	if s.tls {
		return e.SendWithTLS(addr, auth, &tls.Config{ServerName: s.host})
	}
	return e.Send(addr, auth)
}

// ================ 邮箱验证码服务 ================

// EmailCodeService 邮箱验证码服务
type EmailCodeService struct {
	generator CodeGenerator
	store     CodeStore
	sender    *EmailSender
	expire    int
	template  string
}

func NewEmailCodeService(
	generator CodeGenerator,
	store CodeStore,
	sender *EmailSender,
	expireSeconds int,
	template string,
) *EmailCodeService {
	return &EmailCodeService{
		generator: generator,
		store:     store,
		sender:    sender,
		expire:    expireSeconds,
		template:  template,
	}
}

func (s *EmailCodeService) SendCode(toEmail string) error {
	// 生成验证码
	code := s.generator.Generate()

	// 先存储再发送，确保存储成功后才发送
	err := s.store.Set(toEmail, code, s.expire)
	if err != nil {
		return err
	}

	// 使用刚存储的同一个验证码发送邮件
	content := fmt.Sprintf(s.template, code, s.expire/60)
	return s.sender.Send(toEmail, "登录验证码", content)
}

func (s *EmailCodeService) VerifyCode(email, code string) bool {
	stored, err := s.store.Get(email)
	if err != nil || stored != code {
		return false
	}
	s.store.Del(email)
	return true
}

// ================ 便捷函数 ================

type EmailConfig struct {
	From   string
	Nick   string
	Host   string
	Port   int
	Secret string
}

func SendEmailCode(rds *redis.Redis, toEmail string, emailConfig EmailConfig) error {
	service := NewEmailCodeService(
		NewNumberCodeGenerator(6),
		NewRedisCodeStore(rds, "email_code:"),
		NewEmailSender(
			emailConfig.From,
			emailConfig.Nick,
			emailConfig.Host,
			emailConfig.Port,
			emailConfig.Secret,
			true,
		),
		300,
		`<h3>你的验证码是：%s</h3><p>%d分钟内有效，请勿泄露</p>`,
	)
	return service.SendCode(toEmail)
}

func VerifyEmailCode(rds *redis.Redis, email, code string) bool {
	store := NewRedisCodeStore(rds, "email_code:")
	service := NewEmailCodeService(nil, store, nil, 0, "")
	return service.VerifyCode(email, code)
}
