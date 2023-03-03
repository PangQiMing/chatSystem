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

type GroupMembersController interface {
	Insert(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type groupMemberController struct {
	groupMembersService service.GroupMembersService
	jwtService          service.JWTService
}

func NewGroupMembersController(membersService service.GroupMembersService, jwtService service.JWTService) GroupMembersController {
	return &groupMemberController{groupMembersService: membersService, jwtService: jwtService}
}

// Insert 加入群组
// @Summary 加入群组
// @Schemes
// @Description 加入群组
// @Tags 群组模块
// @param Authorization header string false "Authorization"
// @param group_id query string false "group_id"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/group/groupMembers/insert [post]
func (gm *groupMemberController) Insert(ctx *gin.Context) {
	var groupMembersDTO dto.GroupMembersCreateDTO
	groupID := ctx.Param("group_id")
	convertedGroupID, _ := strconv.ParseUint(groupID, 10, 64)
	groupMembersDTO.GroupID = convertedGroupID
	err := ctx.ShouldBind(&groupMembersDTO)
	if err != nil {
		response := helper.BuildErrResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	userID := gm.getUserIdByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		groupMembersDTO.UserID = convertedUserID
	}
	result := gm.groupMembersService.Insert(groupMembersDTO)
	response := helper.BuildResponse(true, "success", result)
	ctx.JSON(http.StatusOK, response)
}

// Delete 退出群组
// @Summary 退出群组
// @Schemes
// @Description 退出群组
// @Tags 群组模块
// @param Authorization header string false "Authorization"
// @param group_id query string false "group_id"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/group/groupMembers [delete]
func (gm *groupMemberController) Delete(ctx *gin.Context) {
	var groupMembers entity.GroupMembers
	groupID := ctx.Query("group_id")
	fmt.Println(groupID)
	convertedGroupID, _ := strconv.ParseUint(groupID, 10, 64)
	groupMembers.GroupID = convertedGroupID
	authHeader := ctx.GetHeader("Authorization")
	UserID := gm.getUserIdByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(UserID, 10, 64)
	if err == nil {
		groupMembers.UserID = convertedUserID
	}
	gm.groupMembersService.Delete(groupMembers)
	response := helper.BuildResponse(true, "Delete a GroupMembers", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (gm *groupMemberController) getUserIdByToken(tokenStr string) string {
	token, err := gm.jwtService.ValidateToken(tokenStr)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
