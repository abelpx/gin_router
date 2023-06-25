package main

import (
	"github.com/gin-gonic/gin"
	"github.com/git-abel/gin_router"
	_ "github.com/git-abel/gin_router/demo/controller"
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
