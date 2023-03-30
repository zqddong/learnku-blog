package config

import "github.com/zqddong/learnku-blog/pkg/config"

func init() {
	config.Add("pagination", config.StrMap{

		// 默认每页条数
		"perpage": config.Env("PERPAGE", 10),

		// URL 中用以分辨多少页的参数
		"url_query": "page",
	})
}
