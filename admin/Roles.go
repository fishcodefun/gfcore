package admin

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func RolesIndex(r *ghttp.Request) {
	r.Response.Write("RolesIndex")
}
