package utils

import (
	"crypto/tls"
	"fmt"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// 发送邮箱验证码
func SendEmailCode(rds *redis.Redis, toEmail string, emailConfig EmailConfig) error {
	// 1. 生成 6 位数字验证码
	code := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)

	// 2. 存入 Redis，5 分钟过期
	err := rds.Setex("email_code:"+toEmail, code, 300)
	if err != nil {
		return err
	}

	// 3. 构造邮件
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", emailConfig.Nick, emailConfig.From)
	e.To = []string{toEmail}
	e.Subject = "登录验证码"
	e.HTML = []byte(fmt.Sprintf(`<h3>你的验证码是：%s</h3><p>5分钟内有效，请勿泄露</p>`, code))

	// 4. 发送
	portStr := strconv.Itoa(emailConfig.Port)
	err = e.SendWithTLS(
		emailConfig.Host+":"+portStr,
		emailConfig.From,
		emailConfig.Secret,
		&tls.Config{ServerName: emailConfig.Host},
	)

	return err
}

// 校验邮箱验证码
func VerifyEmailCode(rds *redis.Redis, email string, code string) bool {
	key := "email_code:" + email
	// 获取并删除（一次性验证）
	val, err := rds.Get(key)
	if err != nil {
		return false
	}
	rds.Del(key)
	return val == code
}

// 邮箱配置结构
type EmailConfig struct {
	From   string
	Nick   string
	Host   string
	Port   int
	Secret string
}
