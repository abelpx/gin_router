# gin_router
### 什么是 gin_router ？
* 你是否对gin 框架中的路由管理的稀烂深感痛苦？接口多了，对路由文件该如何管理无从下手？路由文件多了，需要根据路由找对应的接口翻半天却找不多？
* 那么新的一种路由解决方案它来了，`gin_router` 是一个适用于 gin 框架的自动注册rest-ful风格的路由的插件。它省略了需要维护的路由文件，直接按照你的接口信息来生成接口对应的rest-ful路由，来供客户端进行调用。

### gin_router 的架构图（起名有点高调了 😁）


### gin_router 做了什么呢？ 给个🌰，先看看能否理解
#### 安装
go get github.com/git-abel/gin_router

#### main.go 中的逻辑以及相关引用解释
```go
package main

import (
	_ "controller" // 这里是你 API 的目录，在 main.go 中引用API地址后，会自动对这个目录下的接口进行注册
	"github.com/git-abel/gin_router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	gin_router.BindRoute(r)
	// 启动服务
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
```

#### controller(这个是你项目中放置接口的目录) 下接口注册逻辑大概解释
```go
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/git-abel/gin_router"
)

type HomeController struct {
}

// @Group(v1)
func init() {
	gin_router.RegisterRoute(&HomeController{})
}

// Index
// @Method(GET)
func (c *HomeController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Hi
// @Method(GET)
func (c *HomeController) Hi(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Delete
// @Method(Delete)
func (c *HomeController) Delete(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Create
// @Method(Post)
func (c *HomeController) Create(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Update
// @Method(PUT)
func (c *HomeController) Update(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// MemberName
// @Method(PUT)
// @Member(:names)
func (c *HomeController) MemberName(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("names"))
}

// Member
// @Method(PUT)
// @Member(:ids)
func (c *HomeController) Member(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("ids"))
}

// MoreMember
// @Method(PUT)
// @Member(:ids/eeeee/:names)
func (c *HomeController) MoreMember(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("ids")+"  names "+ctx.Params.ByName("names"))
}

// MoreBMember
// @Method(PUT)
// @Member(:ids/bbb/:names)
func (c *HomeController) MoreBMember(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("ids")+"  names "+ctx.Params.ByName("names"))
}


```

###### 其上输出的路由分别是：
| 请求方法   | 路径                                     |
|--------|----------------------------------------|
| GET    | /v1/home                               |
| DELETE | /v1/home/:id                           |
| PUT    | /v1/home/:id                           |
| GET    | /v1/home/hi                            |
| PUT    | /v1/home/member_name/:names            |
| PUT    | /v1/home/more_bmember/:ids/bbb/:names  |
| POST   | /v1/home                               |
| PUT    | /v1/home/member/:ids                   |
| PUT    | /v1/home/more_member/:ids/eeeee/:names |

###### 在上面的示例代码中 init() 函数是通过调用 `gin_router.RegisterRoute` 函数，来将这个文件中的 Handler 注册进入路由，其中注解的详细解释如下：

| 注解名称    | 注解含义                                                                                                                                              |
|---------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| @Group  | 此处括号中的名称起到一个 namespace 的作用，为路由增加一个访问层级，如在 /home/hi 前增加 v1/home/hi，此处的 @Group 可以省略不写，路由会自动通过 controller 下的目录结构来生成路由地址，可以通过目录结构来添加 v1、v2            |
| @Method | 此处括号中的名称用来规定此接口需要通过什么请求方式来进行访问（建议每个接口规定一种访问方式，不过此处支持多种访问方式用作一个接口请求），目前支持rest-ful中标准的四种访问类型 get、post、delete、put。（*@Method中大小写都支持 get、Get、gEt ...*） |
| @Member | URL 中通配符的内容用于生成 xxx/xxx/:name，其中 :name 就是 @Member 中的内容，可以适配多个通配符 :name/sdf/:age/:uid                                                              |


###### 为什么 Index 接口并未输出为 v1/home/index 呢？
>
> 本插件做了点多余的操作，其 index 与 create 并不会生成后缀。简化其路由信息。


#### 额外配置支持
`[route_config.yml](demo%2Froute_config.yml)` *此配置文件可忽略，此文件若要配置需放置在项目的根目录*

| 字段名称        | 作用                                                         |
|-------------|------------------------------------------------------------|
| omit_suffix | 若你的接口名称中存在统一的后缀名称，如 `Controller` 并且在生成路由时想隐藏它，则添加此配置用作后缀忽略 |
| api_path    | 你放置 Api 的路径如果不放置, 则默认在根目录下查找 controller 目录里面的接口            |


### 下一步计划
拓展 gin.Context, 并将其无感的并入此插件中。

