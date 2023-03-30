package main

import (
	"embed"
	"github.com/zqddong/learnku-blog/app/http/middlewares"
	"github.com/zqddong/learnku-blog/bootstrap"
	"github.com/zqddong/learnku-blog/config"
	c "github.com/zqddong/learnku-blog/pkg/config"
	"github.com/zqddong/learnku-blog/pkg/logger"
	"net/http"
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

//go:embed public/*
var staticFS embed.FS

func init() {
	// 初始化配置信息
	config.Initialize()

}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化模板
	bootstrap.SetupTemplate(tplFS)

	// 初始化路由绑定
	router := bootstrap.SetupRoute(staticFS)

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
