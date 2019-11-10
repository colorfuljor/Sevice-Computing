package service

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/unrolled/render"
)

//Run 运行服务器
func Run(port string) {
	m := martini.Classic()
	m.Use(martini.Static("static"))
	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		r := render.New()
		r.HTML(res, http.StatusOK, "index", "World")
	})
	m.RunOnAddr(":" + port)
}
