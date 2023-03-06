package controller

import (
	"chat/dto"
	"chat/entity"
	"chat/helper"
	"chat/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CircleController interface {
	All(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type circleController struct {
	circleService service.CircleService
	jwtService    service.JWTService
}

func NewCircleController(circleService service.CircleService, jwtService service.JWTService) CircleController {
	return &circleController{circleService: circleService, jwtService: jwtService}
}

// All 获取我的动态
// @Summary 获取我的动态
// @Schemes
// @Description 获取我的动态
// @Tags 朋友圈模块
// @param Authorization header string false "Authorization"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/moment/all [get]
func (b *circleController) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id := b.getUserIdByToken(authHeader)

	momentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := helper.BuildErrResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	moments := b.circleService.All(momentID)
	response := helper.BuildResponse(true, "ok!", moments)
	ctx.JSON(http.StatusOK, response)
}

// Insert 发布我的动态
// @Summary 发布我的动态
// @Schemes
// @Description 发布我的动态
// @Tags 朋友圈模块
// @param Authorization header string false "Authorization"
// @param title query string false "title"
// @param description query string false "description"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/moment/insert [post]
func (b *circleController) Insert(ctx *gin.Context) {
	var circleCreateDTO dto.CircleCreateDTO
	err := ctx.ShouldBind(&circleCreateDTO)
	if err != nil {
		response := helper.BuildErrResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := b.getUserIdByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			circleCreateDTO.UserID = convertedUserID
		}
		result := b.circleService.Insert(circleCreateDTO)
		response := helper.BuildResponse(true, "ok!", result)
		ctx.JSON(http.StatusOK, response)
	}
}

// Delete 删除我的动态
// @Summary 删除我的动态
// @Schemes
// @Description 删除我的动态
// @Tags 朋友圈模块
// @param Authorization header string false "Authorization"
// @param id query string false "moment id"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/moment/delete [delete]
func (b *circleController) Delete(ctx *gin.Context) {
	var moment entity.Circle
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrResponse("Failed to get id", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	}
	moment.ID = id
	authHeader := ctx.GetHeader("Authorization")
	userID := b.getUserIdByToken(authHeader)
	if b.circleService.IsAllowedToDelete(userID, moment.ID) {
		b.circleService.Delete(moment)
		response := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (b *circleController) getUserIdByToken(tokenStr string) string {
	token, err := b.jwtService.ValidateToken(tokenStr)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
