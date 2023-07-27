package gin_router

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type RegisterTest struct {
}

// Index
// @Method(GET)
func (c *RegisterTest) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Hi
// @Method(GET)
func (c *RegisterTest) Hi(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Delete
// @Method(Delete)
func (c *RegisterTest) Delete(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Create
// @Method(Post)
func (c *RegisterTest) Create(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Update
// @Method(PUT)
func (c *RegisterTest) Update(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// MemberName
// @Method(PUT)
// @Member(:names)
func (c *RegisterTest) MemberName(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("names"))
}

// Member
// @Method(PUT)
// @Member(:ids)
func (c *RegisterTest) Member(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("ids"))
}

// MoreMember
// @Method(PUT)
// @Member(:ids/eeeee/:names)
func (c *RegisterTest) MoreMember(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("ids")+"  names "+ctx.Params.ByName("names"))
}

// MoreBMember
// @Method(PUT)
// @Member(:ids/bbb/:names)
func (c *RegisterTest) MoreBMember(ctx *gin.Context) {
	ctx.String(200, "Hello, "+ctx.Params.ByName("ids")+"  names "+ctx.Params.ByName("names"))
}

func TestRegisterRoute(t *testing.T) {
	RegisterRoute(&RegisterTest{})
}
