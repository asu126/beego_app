package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
	//

	"net/http/httputil"
)

type PublicController struct {
	beego.Controller
}

type handle struct {
	host string
	port string
}

func (this *PublicController) Get() {
	rub_mode := beego.AppConfig.String("runmode")

	if rub_mode == "dev" {
		// http://localhost:8081
		webpack_host := beego.AppConfig.String("webpack_host")
		webpack_port := beego.AppConfig.String("webpack_port")

		remote, err := url.Parse("http://" + webpack_host + ":" + webpack_port)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(this.Ctx.ResponseWriter, this.Ctx.Request)
	} else {
		filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}

		f, err := os.Open(filePath)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
		defer f.Close()

		_, err = io.Copy(this.Ctx.ResponseWriter, f)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
	}
}
