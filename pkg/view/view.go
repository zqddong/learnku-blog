package view

import (
	"embed"
	"github.com/zqddong/learnku-blog/app/models/category"
	"github.com/zqddong/learnku-blog/app/models/user"
	"github.com/zqddong/learnku-blog/pkg/auth"
	"github.com/zqddong/learnku-blog/pkg/flash"
	"github.com/zqddong/learnku-blog/pkg/logger"
	"github.com/zqddong/learnku-blog/pkg/route"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

type D map[string]interface{}

var TplFS embed.FS

// Render 渲染通用视图
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单的视图
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {

	// 1. 通用模板数据
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	data["Categories"], _ = category.All()

	// 2. 生成模板文件
	allFiles := getTemplateFiles(tplFiles...)

	// 3. 解析所有模板文件
	//tmpl, err := template.New("").
	//	Funcs(template.FuncMap{
	//		"RouteName2URL": route.Name2URL,
	//	}).ParseFiles(allFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFS(TplFS, allFiles...)
	logger.LogError(err)

	// 4. 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	// 1 设置模板相对路径
	viewDir := "resources/views/"

	// 2. 遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 3. 所有布局模板文件 Slice
	//layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	layoutFiles, err := fs.Glob(TplFS, viewDir+"layouts/*.gohtml")
	logger.LogError(err)

	// 4. 合并所有文件
	return append(layoutFiles, tplFiles...)
}

func One(w io.Writer, data interface{}) {
	tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
	logger.LogError(err)

	// 3. 渲染模板，将所有文章的数据传输进去
	err = tmpl.Execute(w, data)
	logger.LogError(err)
}
