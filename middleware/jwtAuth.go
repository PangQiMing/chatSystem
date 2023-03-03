package middleware

import (
	"chat/helper"
	"chat/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrResponse("处理请求失败", "没有找到Token", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id]: ", claims["user_id"])
			log.Println("Claims[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrResponse("Token验证失败", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}
}
