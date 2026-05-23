// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"fmt"
	"net/http"

	"goZeroApi/internal/config"
	"goZeroApi/internal/handler"
	"goZeroApi/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gozeroapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// ========= 👇 只加这一段，swagger 就集成好了 =========
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/swagger",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./swagger/index.html")
		}),
	})

	// 2. 【关键！！！】 必须加这个路由！！！
	// 因为页面请求的是 /doc/swagger.json
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/doc/swagger.json",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./doc/swagger.json")
		}),
	})
	// ====================================================

	server.Start()
}
