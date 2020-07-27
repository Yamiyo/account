package http

import (
	"github.com/Yamiyo/account/glob"
	_ "github.com/Yamiyo/account/transport/http/controller"

	"github.com/gin-gonic/gin"

	"reflect"
)

func InitController() (*gin.Engine, error) {
	router := gin.New()

	for _, control := range glob.Controllers {
		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(router)
		control.Call(params)
	}

	return router, nil
}
