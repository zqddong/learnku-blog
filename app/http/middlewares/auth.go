package middlewares

import (
	"github.com/zqddong/learnku-blog/pkg/auth"
	"github.com/zqddong/learnku-blog/pkg/flash"
	"net/http"
)

// HttpHandlerFunc 简写 —— func(http.ResponseWriter, *http.Request)
type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

// Auth 登录用户才可访问
// 单路由中间件
func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next(w, r)
	}
}
