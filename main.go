package main

import (
	"github.com/zqddong/learnku-blog/app/http/middlewares"
	"github.com/zqddong/learnku-blog/bootstrap"
	"github.com/zqddong/learnku-blog/config"
	c "github.com/zqddong/learnku-blog/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()

}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
