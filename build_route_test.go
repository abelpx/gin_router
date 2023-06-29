package gin_router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBindRoute(t *testing.T) {
	// Routers 结构体
	Routers = map[string]map[string]Route{
		"v1": {
			"create": Route{
				Path:        "/homes_controller",
				Method:      []string{"Post"},
				Action:      "create",
				HandlerFunc: func(c *gin.Context) {},
			},
			"delete": Route{
				Path:        "/homes_controller",
				Method:      []string{"Delete"},
				Action:      "delete",
				HandlerFunc: func(c *gin.Context) {},
			},
			"hi": Route{
				Path:        "/homes_controller",
				Method:      []string{"GET"},
				Action:      "hi",
				HandlerFunc: func(c *gin.Context) {},
			},
			"update": Route{
				Path:        "/homes_controller",
				Method:      []string{"PUT"},
				Action:      "update",
				HandlerFunc: func(c *gin.Context) {},
			},
			"index": Route{
				Path:        "/homes_controller",
				Method:      []string{"GET", "POST"},
				Action:      "index",
				HandlerFunc: func(c *gin.Context) {},
			},
		},
	}

	fmt.Print(Routers)

	// 创建一个 gin 引擎
	r := gin.Default()

	// 绑定路由
	BindRoute(r)

	// 创建一个 GET 请求
	req, err := http.NewRequest("GET", "/v1/homes_controller", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个响应记录器
	rec := httptest.NewRecorder()

	// 处理请求
	r.ServeHTTP(rec, req)

	// 检查响应状态码是否为 200
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 创建一个 POST 请求
	req, err = http.NewRequest("POST", "/v1/homes_controller", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 重置响应记录器
	rec = httptest.NewRecorder()

	// 处理请求
	r.ServeHTTP(rec, req)

	// 检查响应状态码是否为 200
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/v1/homes_controller/hi", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 重置响应记录器
	rec = httptest.NewRecorder()

	// 处理请求
	r.ServeHTTP(rec, req)

	// 检查响应状态码是否为 200
	if status := rec.Code; status != http.StatusOK {
		fmt.Println("123123123")
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 创建一个 PUT 请求
	req, err = http.NewRequest("PUT", "/v1/homes_controller/update", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 重置响应记录器
	rec = httptest.NewRecorder()

	// 处理请求
	r.ServeHTTP(rec, req)

	// 检查响应状态码是否为 200
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 创建一个 DELETE 请求
	req, err = http.NewRequest("DELETE", "/v1/homes_controller/delete", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 重置响应记录器
	rec = httptest.NewRecorder()

	// 处理请求
	r.ServeHTTP(rec, req)

	// 检查响应状态码是否为 200
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
