package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Yamiyo/account/glob"
	"github.com/Yamiyo/account/service/session"
)

func init() {
	glob.AutoRegister(&User{})
}

type User struct {
	SessionService session.SessionService `injection:"sessionService"`
}

func (user *User) SetupRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1/user")
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)

		v1.DELETE("/logout", user.Logout)
	}
}

func (user *User) Login(ctx *gin.Context) {

}

func (user *User) Logout(ctx *gin.Context) {

}

func (user *User) Register(ctx *gin.Context) {

}
