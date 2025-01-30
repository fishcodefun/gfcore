package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ViewLogin(r *ghttp.Request, sitename string) {
	response := g.Map{
		"SiteName": sitename,
	}
	r.Response.WriteTpl("/view/index.tpl", response)
}

func AuthLogin(r *ghttp.Request, sitename string) {
	r.Response.RedirectTo("/")
	// r.Response.WriteJson(g.Map{
	// 	"errors": "账号或密码错误",
	// 	"data":   "",
	// 	"status": false,
	// })
}
