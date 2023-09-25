package main

import (
	// _ "demo/app/interfaces"
	_ "demo/controller"
	"github.com/gin-gonic/gin"
	"github.com/git-abel/gin_router"
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
