package admin

import (
	"github.com/fishcodefun/gfcore/model"
	"github.com/fishcodefun/gfcore/server"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func UsersIndex(r *ghttp.Request) {
	//查询数据
	var users []model.AdminUsers
	if err := g.DB().Model("admin_users").Scan(&users); err != nil {
		r.Response.Write("查询数据失败: " + err.Error())
		return
	}
	var usersAny []any
	for _, user := range users {
		usersAny = append(usersAny, user)
	}

	// 定义列信息
	columns := []server.Column{
		server.Field("Id", "ID").Sort(),
		server.Field("Username", "账号").Using([]string{"不可售", "可售"}),
		server.Field("Avatar", "头像"),
		server.Field("Name", "姓名"),
		server.Field("role", "角色"),
		server.Field("permissions", "权限"),
		server.Field("CreatedAt", "创建时间"),
		server.Field("UpdatedAt", "更新时间"),
	}

	response := server.Response(usersAny, columns)

	filters := []server.TableFilter{
		{Field: "username", Label: "用户名", Width: 120},
		{Field: "name", Label: "姓名"},
	}
	response["filter"] = filters

	//g.View().BindFunc("hello", funcHello)
	//reflect.TypeOf(v).String()

	// 渲染模板
	r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.Response.WriteTpl("/view/Table.tpl", response)
}
