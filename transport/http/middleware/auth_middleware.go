package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/Yamiyo/account/glob"
)

func init() {
	glob.AutoRegister(&AuthMiddleware{})
}

type AuthMiddleware struct {
	//LoginService login.LoginService `injection:"LoginService"`
}

func (rc *AuthMiddleware) Role(roles ...string) func(ctx *gin.Context) {
	roleHandle := func(ctx *gin.Context) {
		//LoginService 驗證token + 取出資料
		//該使用者存在 session 的權限是否有包含 roles 的任何 字段
		// 有就是驗證成功, 沒有就是失敗
	}
	return roleHandle
}
