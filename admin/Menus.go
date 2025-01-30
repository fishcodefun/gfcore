package admin

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func MenusIndex(r *ghttp.Request) {
	r.Response.Write("MenusIndex")
}
