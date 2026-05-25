package pkg

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// RedisStore Redis存储实现
type RedisStore struct {
	Redis *redis.Redis
}

func NewRedisStore(redis *redis.Redis) *RedisStore {
	return &RedisStore{Redis: redis}
}

// Set 设置验证码
func (s *RedisStore) Set(id string, value string) error {
	// 5分钟过期
	return s.Redis.Setex("captcha:"+id, value, 300)
}

// Get 获取验证码
func (s *RedisStore) Get(id string, clear bool) string {
	key := "captcha:" + id
	logx.Info("captchaGet", key)

	val, _ := s.Redis.Get(key)
	logx.Info("captchaGetVal", val)

	if clear {
		// 验证后删除
		s.Redis.Del(key)
	}

	return val
}

// Verify 验证验证码
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	logx.Info("Verify", id, answer)
	return s.Get(id, clear) == answer
}
