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

type FriendController interface {
	Insert(ctx *gin.Context)
	Delete(ctx *gin.Context)
	AllFriend(ctx *gin.Context)
}

type friendController struct {
	friendService service.FriendService
	jwtService    service.JWTService
}

func NewFriendController(friendService service.FriendService, jwtService service.JWTService) FriendController {
	return &friendController{friendService: friendService, jwtService: jwtService}
}

// Insert 添加好友
// @Summary 添加好友
// @Schemes
// @Description 添加好友
// @Tags 好友模块
// @param Authorization header string false "Authorization"
// @param friend_email query string false "friend_email"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/friend [post]
func (f *friendController) Insert(ctx *gin.Context) {
	var friendDTO dto.FriendDTO
	err := ctx.ShouldBind(&friendDTO)
	if err != nil {
		response := helper.BuildErrResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	userID := f.getUserIdByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		friendDTO.UserID = convertedUserID
	}
	result, res := f.friendService.Insert(friendDTO)
	if res == 1 {
		response := helper.BuildErrResponse("system xxx", "friend already exists", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else if res == -1 {
		response := helper.BuildErrResponse("system xxx", "not found the friend ", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, "success", result)
	ctx.JSON(http.StatusOK, response)
}

// Delete 删除好友
// @Summary 删除好友
// @Schemes
// @Description 删除好友
// @Tags 好友模块
// @param Authorization header string false "Authorization"
// @param friend_email query string false "friend_email"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/friend [delete]
func (f *friendController) Delete(ctx *gin.Context) {
	var friend entity.Friend
	friendEmail := ctx.Param("friend_email")
	friend.FriendEmail = friendEmail
	authHeader := ctx.GetHeader("Authorization")
	userID := f.getUserIdByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		friend.UserID = convertedUserID
	}
	f.friendService.Delete(friend, convertedUserID)
	response := helper.BuildResponse(true, "Delete a friend", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

	//response := helper.BuildErrResponse("invalid operation", "There is not friend in your list", helper.EmptyObj{})
	//ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

// AllFriend 查询所有好友
// @Summary 查询所有好友
// @Schemes
// @Description 查询所有好友
// @Tags 好友模块
// @param Authorization header string false "Authorization"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/friend [get]
func (f *friendController) AllFriend(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	userID := f.getUserIdByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		friend := f.friendService.AllFriend(convertedUserID)
		response := helper.BuildResponse(true, "ok!", friend)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrResponse("invalid operation", "There is not friend in your list", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (f *friendController) getUserIdByToken(tokenStr string) string {
	token, err := f.jwtService.ValidateToken(tokenStr)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
