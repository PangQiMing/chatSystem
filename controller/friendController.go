package controller

import (
	"chat/dto"
	"chat/entity"
	"chat/helper"
	"chat/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type FriendController interface {
	Insert(ctx *gin.Context)
	Delete(ctx *gin.Context)
	AllFriend(ctx *gin.Context)
	FindFriendByEmail(ctx *gin.Context)
	ShowAddFriendList(ctx *gin.Context)
}

type friendController struct {
	friendService service.FriendService
	jwtService    service.JWTService
	userService   service.UserService
}

func NewFriendController(friendService service.FriendService, userService service.UserService, jwtService service.JWTService) FriendController {
	return &friendController{friendService: friendService, userService: userService, jwtService: jwtService}
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
// @Router /api/friend/add [post]
func (f *friendController) Insert(ctx *gin.Context) {
	var addFriendDTO dto.AddFriendDTO
	err := ctx.ShouldBind(&addFriendDTO)
	if err != nil {
		response := helper.BuildErrResponse("处理请求失败...", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	userID := f.getUserIdByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		addFriendDTO.UserID = convertedUserID
	}
	result, res := f.friendService.Insert(addFriendDTO)
	if res == 1 {
		response := helper.BuildErrResponse("该用户已经是你的好友", "friend already exists", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else if res == -1 {
		response := helper.BuildErrResponse("查找不到该用户", "not found the friend ", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, "添加成功", result)
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

func (f *friendController) FindFriendByEmail(ctx *gin.Context) {
	var friendSearch dto.FriendDTO
	err := ctx.ShouldBind(&friendSearch)
	if err != nil {
		log.Println(err)
		response := helper.BuildErrResponse("处理请求失败...", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	userID := f.getUserIdByToken(authHeader)
	_, err = strconv.ParseUint(userID, 10, 64)
	if err != nil {
		response := helper.BuildErrResponse("处理请求失败...", "token错误", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	friend := f.userService.FindByEmail(friendSearch.Email)
	if (friend == entity.User{}) {
		response := helper.BuildErrResponse("查找不到此邮箱", "", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "ok!", friend)
		ctx.JSON(http.StatusOK, response)
	}
}

func (f *friendController) ShowAddFriendList(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	userID := f.getUserIdByToken(authHeader)
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		response := helper.BuildErrResponse("处理请求失败...", "token错误", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	user := f.userService.FindUserByID(id)
	result := f.friendService.ShowAddFriendList(user.Email)
	if len(result) > 0 {
		response := helper.BuildResponse(true, "ok!", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrResponse("当前没有新朋友", "", helper.EmptyObj{})
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
