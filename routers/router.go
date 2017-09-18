package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	apiPattern        = `^/api/`
	ciAPIPattern      = `^/ci/api/`
	gitProjectPattern = `^/([^/]+/){1,}[^/]+\.git/`
	projectPattern    = `^/([^/]+/){1,}[^/]+/`
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// gitProjectPattern = `^/([^/]+/){1,}[^/]+\.git/`
	// beego.Get("/([^/]+/){1,}[^/]+.git", func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("hello world"))
	// })
	// beego.Router("/api/:id([0-9]+)", func(ctx *context.Context) {
	// 	// beego.Router("/user/:username([\w]+",func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("hello world dd"))
	// })
	beego.Router("/public/:all", &controllers.PublicController{})
	beego.Router("/help", &controllers.HelpController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/dashbord", &controllers.DashbordController{})

	ns :=
		beego.NewNamespace("/dashbord",
			// beego.NSCond(func(ctx *context.Context) bool {
			// 	if ctx.Input.Domain() == "api.beego.me" {
			// 		return true
			// 	}
			// 	return false
			// }),
			// beego.NSBefore(auth),
			beego.NSGet("/notallowed", func(ctx *context.Context) {
				ctx.Output.Body([]byte("notAllowed"))
			}),
			beego.NSRouter("/reports", &controllers.DashbordController{}, "get:Reports"),
			beego.NSNamespace("/cms",
				beego.NSInclude(
					&controllers.MainController{},
				),
			),
		)
	beego.AddNamespace(ns)

	api_ns :=
		beego.NewNamespace("/api/v1",
			beego.NSNamespace("/shop",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("shopinfo"))
				}),
			),
			beego.NSNamespace("/order",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("orderinfo"))
				}),
			),
			beego.NSNamespace("/crm",
				beego.NSGet("/:id", func(ctx *context.Context) {
					ctx.Output.Body([]byte("crminfo"))
				}),
			),
		)
	beego.AddNamespace(api_ns)

	beego.Router("/homepage", &controllers.HomepageController{})
	beego.Router("/vue", &controllers.VueController{})
	beego.Router("/user/:username([0-9a-zA-Z_-]+)", &controllers.HomepageController{})
	// beego.Get("/api", func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("hello api"))
	// })
}
