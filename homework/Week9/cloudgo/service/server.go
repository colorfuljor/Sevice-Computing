package service

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/unrolled/render"
)

//Run 运行服务器
func Run(port string) {
	m := martini.Classic()

	//加载静态文件资源
	m.Use(martini.Static("static"))

	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		r := render.New()
		//渲染HTML文件
		r.HTML(res, http.StatusOK, "index", "World")
	})
	//使用port端口运行
	m.RunOnAddr(":" + port)
}
