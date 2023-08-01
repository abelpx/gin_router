package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/git-abel/gin_router"
)

type InterController struct {
}

// @Group(v1)
func init() {
	gin_router.RegisterRoute(&InterController{})
}

// Index
// @Method(GET, POST)
func (c *InterController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, inter!")
}
