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
