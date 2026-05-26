package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 统一返回体
type Response struct {
	Code int         `json:"code"` // 错误码 0=成功 非0=失败
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据
}

// 成功返回
func Success(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// 失败返回
func Fail(w http.ResponseWriter, code int, msg string) {
	httpx.OkJson(w, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
