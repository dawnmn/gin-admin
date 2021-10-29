package app

import (
	"gin-admin/router"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func InitRouter() {
	Router = gin.New()
	Router.Use(LogMiddle, gin.Recovery())

	re := reflect.ValueOf(router.Routers{})
	for i := 0; i < re.NumField(); i++ {
		re := re.Field(i)
		t := re.Type()
		for i := 0; i < re.NumMethod(); i++ {
			re := re
			method := t.Method(i)
			api := strings.Trim(t.String(), "*router.") + "/" + method.Name
			Router.POST(api, func(c *gin.Context) {
				params := []reflect.Value{reflect.ValueOf(c)}
				re.MethodByName(method.Name).Call(params)
			})
		}
	}
}
