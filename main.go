package main

import (
	"chat/config"
	"chat/controller"
	"chat/docs"
	"chat/middleware"
	"chat/repository"
	"chat/service"
	"chat/service/ws"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// DB database connection
	DB = config.SetupDatabaseConnection()

	//Repository
	userRepository         = repository.NewUserConnection(DB)
	momentRepository       = repository.NewMomentRepository(DB)
	friendRepository       = repository.NewFriendRepository(DB)
	groupRepository        = repository.NewGroupRepository(DB)
	groupMembersRepository = repository.NewGroupMembersRepository(DB)

	//Service
	authService         = service.NewAuthService(userRepository)
	jwtService          = service.NewJWTService()
	userService         = service.NewUserService(userRepository)
	momentService       = service.NewBookService(momentRepository)
	friendService       = service.NewFriendService(friendRepository)
	groupService        = service.NewGroupService(groupRepository)
	groupMembersService = service.NewGroupMembersService(groupMembersRepository)

	//Controller
	authController         = controller.NewAuthController(authService, jwtService)
	userController         = controller.NewUserController(userService, jwtService)
	momentController       = controller.NewMomentController(momentService, jwtService)
	friendController       = controller.NewFriendController(friendService, jwtService)
	groupController        = controller.NewGroupController(groupService, jwtService)
	groupMembersController = controller.NewGroupMembersController(groupMembersService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(DB)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ws
	hub := ws.NewHub()
	go hub.Run()
	authRouters := r.Group("api/auth")
	{
		authRouters.POST("login", authController.Login)
		authRouters.POST("register", authController.Register)
	}
	userRouters := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRouters.GET("profile", userController.Profile)
		userRouters.PUT("profile", userController.Update)
	}

	momentRouters := r.Group("api/moment", middleware.AuthorizeJWT(jwtService))
	{
		momentRouters.POST("/insert", momentController.Insert)
		momentRouters.DELETE("/delete", momentController.Delete)
		momentRouters.PUT("/update", momentController.Update)
		momentRouters.GET("/all", momentController.All)
		momentRouters.GET("/find", momentController.FindByID)
	}

	friendRouters := r.Group("api/friend", middleware.AuthorizeJWT(jwtService))
	{
		friendRouters.GET("/", friendController.AllFriend)
		friendRouters.POST("/", friendController.Insert)
		friendRouters.DELETE("/:friend_email", friendController.Delete)
	}

	groupRouters := r.Group("api/group", middleware.AuthorizeJWT(jwtService))
	{
		groupRouters.POST("/insert", groupController.Insert)
		groupRouters.DELETE("/delete", groupController.Delete)
		groupRouters.PUT("/update", groupController.Update)
		groupRouters.GET("/groupsIManage", groupController.GroupsIManage)

		groupRouters.POST("/groupMembers/insert", groupMembersController.Insert)
		groupRouters.DELETE("/groupMembers", groupMembersController.Delete)
	}

	//msgRouters := r.Group("api/message", middleware.AuthorizeJWT(jwtService))
	msgRouters := r.Group("api/message")
	{
		msgRouters.GET("/ws", func(ctx *gin.Context) {
			ws.ServeWs(hub, ctx.Writer, ctx.Request)
		})

	}
	err := r.Run(":8081")
	if err != nil {
		panic("Router start failed")
	}
}
