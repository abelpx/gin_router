package gin_router

type TestController struct {
}

// @Group(namespace)
func init() {
	RegisterRoute(&TestController{})
}

// func Index
