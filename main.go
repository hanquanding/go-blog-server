/**
 * @author: hqd
 * @description: main
 * @file: main
 * @date: 2021-02-02 23:11
 */

package main

import (
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/hqd8080/go-blog-server/routers"
	"github.com/hqd8080/go-blog-server/services"
)

func init() {
	web.ErrorHandler("404", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"ret":404, "msg":"API not found"}`))
		writer.WriteHeader(404)
	})
}

func init() {
	if err := services.InitLogs(); err != nil {
		panic(err)
	}
	if err := services.InitRedis(); err != nil {
		panic(err)
	}
	if err := services.InitDB(); err != nil {
		panic(err)
	}
}

func main() {
	if web.BConfig.RunMode == web.DEV {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	addr, _ := web.AppConfig.String("httpaddr")
	port, _ := web.AppConfig.String("httpport")
	logs.Info("server start at %s:%s", addr, port)

	web.Run()
}
