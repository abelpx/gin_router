package gin_router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go/ast"
	"go/parser"
	"go/token"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

type Route struct {
	Method      []string        `desc:"RESTful 路由请求方式 [get post put delete ...]"`
	HandlerFunc gin.HandlerFunc `desc:"资源名称"`
	Path        string          `desc:"资源路径"`
	Member      string          `desc:"通配符成员信息"`
	Action      string          `desc:"函数名称"`
}

type Config struct {
	OmitSuffix string `yaml:"omit_suffix" desc:"API 后缀（如果存在统一需要忽略的后缀则通过此方式忽略，如：HelloController{} 如果地址不想携带 Controller , 则将 Controller 添加此处）"`
	ApiPath    string `yaml:"api_path" desc:"放置API的目录地址，注: 需要从项目根目录开始"`
}

// Routers {
//   Group: {
//     ActionName: {
//       Method: [get, Post, PUT],
//       Action: func(*gin.Context)
//     }
//   }
// }
var Routers = make(map[string]map[string]Route)
var ConfigInfo = new(Config)

func init() {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/route_config.yml")
	if err != nil {
		fmt.Println("使用默认路由配置进行构建路由", err.Error())
	} else {
		// 读取配置文件
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&ConfigInfo)
		if err != nil {
			fmt.Println("路由配置文件不合法，自定义配置使用失败，将采用默认进行构建 ... ")
		}
	}
	defer file.Close()
}

func RegisterRoute(controller any) {
	// 临时声明一个 group 字段，方便真正便利到 group 的时候来进行替换
	var tmpGroup string
	// 每此注册，先临时声明一个保存此次 function 的结果
	tmpRoute := make(map[string]Route)

	v := reflect.TypeOf(controller)
	if v.NumMethod() == 0 {
		return
	}
	// 获取 controller 名称
	controllerName := v.Elem().Name()
	// 获取 controller 路径
	controllerPath := ""
	notProjectPath := ""
	if ConfigInfo.ApiPath == "" {
		controllerPath = v.Elem().PkgPath()
		notProjectPath = strings.TrimPrefix(controllerPath, filepath.Dir(controllerPath)+"/")
	} else {
		controllerPath = ConfigInfo.ApiPath
		notProjectPath = ""
	}

	var goFilePath string
	if ConfigInfo.ApiPath == "" {
		goFilePath = filepath.Join(notProjectPath, ToSnakeCase(controllerName)+".go")
	} else {
		goFilePath = filepath.Join(ConfigInfo.ApiPath, ToSnakeCase(controllerName)+".go")
	}
	// 将 MyController => my
	controllerShortName := ToSnakeCase(strings.TrimSuffix(controllerName, ConfigInfo.OmitSuffix))
	// controller 文件夹里面的路径
	path := filepath.Join(notProjectPath, controllerShortName)
	if ConfigInfo.ApiPath == "" {
		path = strings.TrimPrefix(path, "controller")
	} else {
		path = strings.TrimPrefix(path, ConfigInfo.ApiPath)
	}

	// 解析源代码文件
	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, goFilePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, decl := range parsedFile.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if fn.Doc == nil {
			continue
		}

		// 查找带有注释的函数
		// for _, comment := range fn.Doc.List {
		//
		// }

		comment := fn.Doc.Text()
		if len(comment) < 1 {
			continue
		}

		groupPattern := regexp.MustCompile(`@Group\((.+)\)`)
		groupMatches := groupPattern.FindStringSubmatch(comment)
		// 当此字段大于 0 说明遍历到 group
		if len(groupMatches) > 0 && fn.Name.Name == "init" {
			// 如果不是 init() 方法上的 @Group 则忽略
			tmpGroup = groupMatches[1]
			continue
		}

		isMemberPattern := regexp.MustCompile(`(@Member)(\((.+)\))?`)
		isMember := isMemberPattern.FindStringSubmatch(comment)
		member := ""
		if len(isMember) >= 1 {
			if len(isMember[3]) > 0 {
				member += isMember[3] + ""
			} else {
				member += ":id"
			}
		}

		routePattern := regexp.MustCompile(`@Method\((.+)\)`)
		matches := routePattern.FindStringSubmatch(comment)
		if len(matches) < 1 {
			continue
		}

		var methods []string
		// 判断是否有多种请求方式，如果有多种请求方式，会根据 "，|," 来进行补充
		if strings.Contains(matches[1], ",") || strings.Contains(matches[1], "，") {
			routeParts := strings.Split(matches[1], ",")
			methods = routeParts
		} else {
			methods = append(methods, matches[1])
		}
		// 获取当前注释的函数名称
		actionName := ToSnakeCase(fn.Name.Name)
		// 没有经过转化的原始 func 名称，方便后续进行路由注册是进行方法查找
		action := fn.Name.Name

		handlerValue := reflect.ValueOf(controller).MethodByName(action)
		if !handlerValue.IsValid() {
			panic(fmt.Sprintf("Action %s not found in controller", action))
		}

		tmpRoute[path+member+actionName] = Route{
			Method:      methods,
			HandlerFunc: createHandlerFunc(handlerValue),
			Path:        path,
			Member:      member,
			Action:      actionName,
		}
		// fmt.Println("path + member", path+member)
	}

	for key, route := range tmpRoute {
		// 当存在 Group 时，将 Group 追加到请求路径中, 不存在的话默认为空
		groupName := strings.Join([]string{"", tmpGroup}, "/")

		if Routers[groupName] == nil {
			Routers[groupName] = make(map[string]Route)
		}

		Routers[groupName][key] = route
	}
}

func createHandlerFunc(handlerValue reflect.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		f, ok := handlerValue.Interface().(func(*gin.Context))
		if !ok {
			panic("Invalid handler function")
		}
		f(c)
	}
}

// ToSnakeCase 驼峰转蛇形
func ToSnakeCase(camelCase string) string {
	// 使用正则表达式将大写字母转换为下划线加小写字母的形式
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snakeCase := re.ReplaceAllString(camelCase, "${1}_${2}")

	// 将字符串中的所有字符转换为小写字母
	snakeCase = strings.ToLower(snakeCase)

	return snakeCase
}
