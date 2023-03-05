package controller

import (
	"chat/dto"
	"chat/entity"
	"chat/helper"
	"chat/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login 用户登录
// @Summary 用户登录
// @Schemes
// @Description 用户登录
// @Tags 用户认证模块
// @param email query string false "邮箱"
// @param password query string false "密码"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrResponse("处理请求错误...", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GeneratedToken(strconv.FormatUint(v.UserId, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "登录成功", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrResponse("邮箱密码错误", "无效的邮箱", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

// Register 用户注册
// @Summary 用户注册
// @Schemes
// @Description 用户注册
// @Tags 用户认证模块
// @param name query string false "名字"
// @param email query string false "邮箱"
// @param password query string false "密码"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrResponse("处理请求失败", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrResponse("处理请求失败", "该邮箱已被注册", helper.EmptyObj{})
		//ctx.JSON(http.StatusConflict, response)
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		createUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GeneratedToken(strconv.FormatUint(createUser.UserId, 10))
		createUser.Token = token
		response := helper.BuildResponse(true, "注册成功，跳转到登录页面", createUser)
		ctx.JSON(http.StatusOK, response)
	}
}
