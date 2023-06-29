package gin_router

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// BindRoute 绑定到 gin 路由中
func BindRoute(r *gin.Engine) {
	// Group
	for groupName, route := range Routers {
		if len(groupName) > 0 {
			group := r.Group(groupName)
			for _, method := range route {
				switch method.Action {
				case "index":
					group.GET(method.Path, method.HandlerFunc)
				case "create":
					group.POST(method.Path, method.HandlerFunc)
				case "delete":
					group.DELETE(method.Path+"/:id", method.HandlerFunc)
				case "update":
					group.PUT(method.Path+"/:id", method.HandlerFunc)
				default:
					methodStr := strings.Join(method.Method, "")

					commonParameters := func() (string, gin.HandlerFunc) {
						if len(method.Member) > 0 {
							return strings.Join([]string{method.Path, method.Action, method.Member}, "/"), method.HandlerFunc
						} else {
							return strings.Join([]string{method.Path, method.Action}, "/"), method.HandlerFunc
						}
					}

					// 1. GET：用于获取资源，可以理解为读取操作
					if strings.Contains(strings.ToLower(methodStr), "get") {
						group.GET(commonParameters())
					}

					// 2. POST：用于创建资源，可以理解为写入操作。
					if strings.Contains(strings.ToLower(methodStr), "post") {
						group.POST(commonParameters())
					}

					// 3. PUT：用于更新资源，可以理解为修改操作。
					if strings.Contains(strings.ToLower(methodStr), "put") {
						group.PUT(commonParameters())
					}

					// 4. DELETE：用于删除资源，可以理解为删除操作
					if strings.Contains(strings.ToLower(methodStr), "delete") {
						group.DELETE(commonParameters())
					}
				}
			}
		}
	}
}
