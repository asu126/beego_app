package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
)

const (
	apiPattern        = `^/api/`
	ciAPIPattern      = `^/ci/api/`
	gitProjectPattern = `^/([^/]+/){1,}[^/]+\.git/`
	projectPattern    = `^/([^/]+/){1,}[^/]+/`
)

// route("GET", gitProjectPattern+`info/refs\z`, git.GetInfoRefsHandler(api)),
// route("POST", gitProjectPattern+`git-upload-pack\z`, contentEncodingHandler(git.UploadPack(api)), isContentType("application/x-git-upload-pack-request")),
// route("POST", gitProjectPattern+`git-receive-pack\z`, contentEncodingHandler(git.ReceivePack(api)), isContentType("application/x-git-receive-pack-request")),
// route("PUT", gitProjectPattern+`gitlab-lfs/objects/([0-9a-f]{64})/([0-9]+)\z`, lfs.PutStore(api, proxy), isContentType("application/octet-stream")),

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
	beego.Router("/homepage", &controllers.HomepageController{})
	beego.Router("/vue", &controllers.VueController{})
	// beego.Get("/api", func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("hello api"))
	// })
}
