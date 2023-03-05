package controller

import (
	"chat/dto"
	"chat/helper"
	"chat/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type UserController interface {
	ModifyProfile(ctx *gin.Context)
	Profile(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{userService: userService, jwtService: jwtService}
}

// ModifyProfile 更新用户信息
// @Summary 更新用户信息
// @Schemes
// @Description 更新用户信息
// @Tags 用户模块
// @param Authorization header string false "Authorization"
// @param name query string false "名字"
// @param email query string false "邮箱"
// @param password query string false "密码"
// @param sex query string false "性别"
// @param age query int false "年龄"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/user/profile [post]
func (u *userController) ModifyProfile(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrResponse("处理请求错误...", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := u.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	//上传头像图片
	file, _ := ctx.FormFile("file")
	name := ctx.PostForm("avatar")
	filename := name + ".jpg"
	if err := ctx.SaveUploadedFile(file, "./static"+filename); err != nil {
		response := helper.BuildErrResponse("图片上传失败...", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userUpdateDTO.UserId = id
	userUpdateDTO.Avatar = "./static" + filename
	user := u.userService.Update(userUpdateDTO)
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
func (u *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := u.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := u.userService.Profile(id)
	response := helper.BuildResponse(true, "ok!", user)
	ctx.JSON(http.StatusOK, response)
}

func (u *userController) ChangePassword(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	newPassword := ctx.PostForm("new_password")
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := u.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	user := u.userService.FindByEmail(email)
	if comparePassword(user.Password, []byte(password)) {
		user.UserId = id
		user.Password = newPassword
		result := u.userService.ChangePass(user)
		response := helper.BuildResponse(true, "修改密码成功", result)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.BuildErrResponse("用户密码不正确", "用户密码不正确", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
