package controller

import (
	"chat/dto"
	"chat/helper"
	"chat/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{userService: userService, jwtService: jwtService}
}

// Update 更新用户信息
// @Summary 更新用户信息
// @Schemes
// @Description 更新用户信息
// @Tags 用户模块
// @param Authorization header string false "Authorization"
// @param name query string false "名字"
// @param email query string false "邮箱"
// @param password query string false "密码"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/user/profile [put]
func (c *userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.UserId = id
	user := c.userService.Update(userUpdateDTO)
	response := helper.BuildResponse(true, "ok!", user)
	ctx.JSON(http.StatusOK, response)
}

// Profile 获取用户信息
// @Summary 获取用户信息
// @Schemes
// @Description 获取用户信息
// @Tags 用户模块
// @param Authorization header string false "Authorization"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/user/profile [get]
func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	response := helper.BuildResponse(true, "ok!", user)
	ctx.JSON(http.StatusOK, response)
}
