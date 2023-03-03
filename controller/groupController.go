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

type GroupController interface {
	Insert(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	GroupsIManage(ctx *gin.Context)
}

type groupController struct {
	groupService service.GroupService
	jwtService   service.JWTService
}

func NewGroupController(groupService service.GroupService, jwtService service.JWTService) GroupController {
	return &groupController{groupService: groupService, jwtService: jwtService}
}

// Insert 创建群组
// @Summary 创建群组
// @Schemes
// @Description 创建群组
// @Tags 群组模块
// @param Authorization header string false "Authorization"
// @param group_name query string false "group_name"
// @param notice query string false "notice"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/group/insert [post]
func (g *groupController) Insert(ctx *gin.Context) {
	var groupCreateDTO dto.GroupCreateDTO
	err := ctx.ShouldBind(&groupCreateDTO)
	if err != nil {
		response := helper.BuildErrResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		groupLeaderID := g.getUserIdByToken(authHeader)
		convertedGroupLeaderID, err := strconv.ParseUint(groupLeaderID, 10, 64)
		if err == nil {
			groupCreateDTO.GroupLeaderID = convertedGroupLeaderID
		}
		result := g.groupService.Insert(groupCreateDTO)
		response := helper.BuildResponse(true, "ok!", result)
		ctx.JSON(http.StatusOK, response)
	}
}

// Delete 解散群组
// @Summary 解散群组
// @Schemes
// @Description 解散群组
// @Tags 群组模块
// @param Authorization header string false "Authorization"
// @param group_id query string false "group_id"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/group/delete [delete]
func (g *groupController) Delete(ctx *gin.Context) {
	//groupID, err := strconv.ParseUint(ctx.Param("group_id"), 0, 0)

	groupID, err := strconv.ParseUint(ctx.Query("group_id"), 0, 0)
	if err != nil {
		response := helper.BuildErrResponse("Failed to get id", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	groupLeaderID := g.getUserIdByToken(authHeader)
	convertedGroupLeaderID, _ := strconv.ParseUint(groupLeaderID, 10, 64)
	if g.groupService.Delete(groupID, convertedGroupLeaderID) {
		response := helper.BuildResponse(true, "Delete Success", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// Update 更新群组信息
// @Summary 更新群组信息
// @Schemes
// @Description 更新群组信息
// @Tags 群组模块
// @param Authorization header string false "Authorization"
// @param id query string false "id"
// @param group_name query string false "group_name"
// @param notice query string false "notice"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/group/update [put]
func (g *groupController) Update(ctx *gin.Context) {
	var groupUpdateDTO dto.GroupUpdateDTO
	//id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	//if err != nil {
	//	response := helper.BuildErrResponse("No param id was found", err.Error(), helper.EmptyObj{})
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	//	return
	//}
	//groupUpdateDTO.ID = id

	err := ctx.ShouldBind(&groupUpdateDTO)
	if err != nil {
		response := helper.BuildErrResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusOK, response)
	}
	authHeader := ctx.GetHeader("Authorization")
	groupLeaderID := g.getUserIdByToken(authHeader)
	convertedGroupLeaderID, err := strconv.ParseUint(groupLeaderID, 10, 64)
	if err == nil {
		groupUpdateDTO.GroupLeaderID = convertedGroupLeaderID
	}
	result, err := g.groupService.Update(groupUpdateDTO)
	if err != nil {
		response := helper.BuildErrResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "Update Success", result)
		ctx.JSON(http.StatusOK, response)
	}
}

// GroupsIManage 我管理的群组
// @Summary 我管理的群组
// @Schemes
// @Description 我管理的群组
// @Tags 群组模块
// @param Authorization header string false "Authorization"
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /api/group/groupsIManage [get]
func (g *groupController) GroupsIManage(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	groupLeaderID := g.getUserIdByToken(authHeader)
	convertedGroupLeaderID, err := strconv.ParseUint(groupLeaderID, 10, 64)
	if err == nil {
		result := g.groupService.MyGroup(convertedGroupLeaderID)
		response := helper.BuildResponse(true, "Group I Manage", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (g *groupController) getUserIdByToken(tokenStr string) string {
	token, err := g.jwtService.ValidateToken(tokenStr)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
