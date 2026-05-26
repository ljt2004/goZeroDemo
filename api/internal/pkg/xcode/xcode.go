package xcoder

// 全局状态码定义
const (
	// 成功
	Success = 0

	// 客户端错误
	BadRequest   = 400 // 参数错误
	Unauthorized = 401 // 未登录
	Forbidden    = 403 // 无权限
	NotFound     = 404 // 资源不存在

	// 服务端错误
	ServerError = 500 // 服务器错误

	// 业务错误（自定义）
	ErrorEmailEmpty    = 1001 // 邮箱不能为空
	ErrorEmailCode     = 1002 // 邮箱验证码错误
	ErrorUserExists    = 1003 // 用户已存在
	ErrorUserNotExists = 1004 // 用户不存在
	ErrorPassword      = 1005 // 密码错误
)

// 状态码对应文案
var msg = map[int]string{
	Success:         "成功",
	BadRequest:      "参数错误",
	Unauthorized:    "请先登录",
	Forbidden:       "无权限访问",
	NotFound:        "资源不存在",
	ServerError:     "服务器异常",
	ErrorEmailEmpty: "邮箱不能为空",
	ErrorEmailCode:  "验证码错误",
}

// 获取错误信息
func Msg(code int) string {
	message, ok := msg[code]
	if !ok {
		return "未知错误"
	}
	return message
}
