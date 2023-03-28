package main

import (
	"github.com/zqddong/learnku-blog/app/http/middlewares"
	"github.com/zqddong/learnku-blog/bootstrap"
	"github.com/zqddong/learnku-blog/pkg/logger"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
