package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization header
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//获取有效部分
		//fmt.Println(tokenString)
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后获取claim中的userid
		userId := claims.UserID
		db := common.InitDB()
		var user model.User
		db.First(&user, userId)

		//判断用户是否存在
		if userId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//用户存在，将信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
