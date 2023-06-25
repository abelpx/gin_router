package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/git-abel/gin_router"
)

type HelloController struct {
}

// @Group(v1)
func init() {
	gin_router.RegisterRoute(&HelloController{})
}

// Index
// @Method(GET, POST)
func (c *HelloController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Hi
// @Method(GET)
func (c *HelloController) Hi(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}
