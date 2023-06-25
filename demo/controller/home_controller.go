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
// @Method(GET, POST)
func (c *HomeController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}

// Hi
// @Method(GET)
func (c *HomeController) Hi(ctx *gin.Context) {
	ctx.String(200, "Hello, world!")
}
