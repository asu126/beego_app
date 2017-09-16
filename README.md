### 简介
这个基于beego框架的app,旨在熟悉beego框架下mvc的使用，以及go语言几web知识的学习


### js
- npm init
- webpack
```
npm install --save webpack
```
- 打包编译
```
node_modules/.bin/webpack
```

### Webpack的强大功能
- 生成Source Maps（使调试更容易）
- 使用webpack构建本地服务器
```
npm install --save-dev webpack-dev-server
```

- Loaders

- css
```
npm install --save style-loader css-loader
```

** webpack 的热加载是通过起一个服务来监控文件的变化，并且是以内存文件的形式生成到内存中，所以beego demo 需要通过内部proxy来实现修改即可见。**

### vue
- install
```
npm install --save vue
```
### cache
- https://beego.me/docs/module/cache.md
- https://godoc.org/github.com/garyburd/redigo/redis

### aession 使用问题（https://github.com/astaxie/beego/issues/23）
beego.BConfig.WebConfig.Session.SessionOn = true
- 在Controller中操作Session
 ```
 import (
     "github.com/astaxie/beego"
)
func (c *AdminController) Logindo() {
    username := c.GetString("username")
    c.SetSession("Adminname", username)        //设置Session
    adminName := c.GetSession("Adminname")     //读取Session
    c.DelSession("Adminname")                  //删除Session
    c.TplNames = "admin/login.html"
    }
}
 ```
- 在非Controller中操作Session
 ```
 import (
    "github.com/astaxie/beego/context"
)
var FilterAdmin = func(ctx *context.Context) {
ctx.Input.CruSession.Get("Adminname")            //读取Session
ctx.Input.CruSession.Set("Adminname","value")    //设置Session
ctx.Input.CruSession.Delete("Adminname")         //删除Session
}
 ```

### 参考说明
[beego]: https://beego.me/docs/intro/
[beego 工程]: https://github.com/beego/admin.
[webpack]: http://www.jianshu.com/p/42e11515c10f
[webpack-1]: https://webpack.js.org/configuration/dev-server/
[vue]: https://cn.vuejs.org/
[vue-demo]: https://segmentfault.com/a/1190000008678236
