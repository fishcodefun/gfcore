package gfcore

import (
	"github.com/fishcodefun/gfcore/admin"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Server = g.Server()
)

const (
	some = "v0.0.1"
)

func AdminAuthHandler(r *ghttp.Request) {
	// 跨域
	r.Response.CORSDefault()
	// 执行路由回调函数
	r.Middleware.Next()
	// 判断是否产生错误
	if err := r.GetError(); err != nil {
		r.Response.Write("error occurs: ", err.Error())
		return
	}
}

type AdminConfig struct {
	Prefix   string `yaml:"prefix"`
	SiteName string `yaml:"siteName"`
}

// 后台路由节点
func Routers(s *ghttp.Server, config AdminConfig) {
	s.Group("/"+config.Prefix, func(group *ghttp.RouterGroup) {
		group.Middleware(AdminAuthHandler)
		group.GET("/auth/login", func(r *ghttp.Request) {
			admin.ViewLogin(r, config.SiteName)
		})
		group.POST("/auth/login", func(r *ghttp.Request) {
			admin.AuthLogin(r, config.SiteName)
		})
		group.GET("/users/", func(r *ghttp.Request) {
			admin.UsersIndex(r)
		})
		group.GET("/menus/", func(r *ghttp.Request) {
			admin.MenusIndex(r)
		})
	})
}

func Resource(group *ghttp.RouterGroup, path string, controller interface{}) {
	subGroup := group.Group(path)
	subGroup.Map(map[string]interface{}{
		"GET:    /":         controller.(interface{ Index(*ghttp.Request) }).Index,
		"GET:    /:id":      controller.(interface{ Show(*ghttp.Request) }).Show,
		"GET:    /create":   controller.(interface{ Create(*ghttp.Request) }).Create,
		"POST:   /":         controller.(interface{ Store(*ghttp.Request) }).Store,
		"GET:    /:id/edit": controller.(interface{ Edit(*ghttp.Request) }).Edit,
		"PUT:    /:id":      controller.(interface{ Update(*ghttp.Request) }).Update,
		"DELETE: /:id":      controller.(interface{ Destroy(*ghttp.Request) }).Destroy,
	})
}

type ResourceController interface {
	Index(r *ghttp.Request)
	Show(r *ghttp.Request)
}

type BaseController struct{}

func (c *BaseController) Index(r *ghttp.Request) {
	r.Response.Write("Base List" + r.URL.Path)
}

func (c *BaseController) Show(r *ghttp.Request) {
	id := r.Get("id")
	r.Response.Write("Show Resource from BaseController: " + id.String())
}

func (c *BaseController) Create(r *ghttp.Request) {
	r.Response.Write("Create")
}

func (c *BaseController) Store(r *ghttp.Request) {
	r.Response.Write("Store")
}

func (c *BaseController) Edit(r *ghttp.Request) {
	r.Response.Write("Edit")
}

func (c *BaseController) Update(r *ghttp.Request) {
	r.Response.Write("Update")
}

func (c *BaseController) Destroy(r *ghttp.Request) {
	r.Response.Write("Destroy")
}
