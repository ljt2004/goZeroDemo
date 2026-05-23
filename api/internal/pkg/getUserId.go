package pkg

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
)

// GetUserIdFromCtx 从 context 中获取 userId（通用、安全、支持雪花ID）
func GetUserIdFromCtx(ctx context.Context) (int64, error) {
	// 从 ctx 中取出值
	val := ctx.Value("userId")
	if val == nil {
		return 0, errors.New("未获取到用户信息，请登录")
	}

	// 自动判断类型，兼容 string / float64 / int64 / json.Number
	switch v := val.(type) {
	case string:
		// 雪花ID最终方案：string → int64
		userId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.New("用户ID格式错误")
		}
		return userId, nil

	case float64:
		// 兼容旧的 JWT（不推荐，大数字会丢精度）
		return int64(v), nil

	case int64:
		// 标准类型
		return v, nil

	default:
		// 尝试处理 json.Number 类型
		if num, ok := val.(interface{ String() string }); ok {
			userId, err := strconv.ParseInt(num.String(), 10, 64)
			if err != nil {
				return 0, errors.Errorf("用户ID格式错误: %v", err)
			}
			return userId, nil
		}
		return 0, errors.Errorf("不支持的userId类型: %T", val)
	}
}
