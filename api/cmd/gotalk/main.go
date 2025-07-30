package main

import (
	"errors"
	"fmt"
	"github.com/ebubekir/go-talk/api/cmd/gotalk/docs"
	"github.com/ebubekir/go-talk/api/internal/middleware"
	"github.com/ebubekir/go-talk/api/internal/response"
	roomApp "github.com/ebubekir/go-talk/api/internal/room/application"
	roomInfra "github.com/ebubekir/go-talk/api/internal/room/infra"
	roomHttp "github.com/ebubekir/go-talk/api/internal/room/interfaces/http"
	userApp "github.com/ebubekir/go-talk/api/internal/user/application"
	userInfra "github.com/ebubekir/go-talk/api/internal/user/infra"
	userHttp "github.com/ebubekir/go-talk/api/internal/user/interfaces/http"
	"github.com/ebubekir/go-talk/api/internal/websocket"
	"github.com/ebubekir/go-talk/api/pkg/firebase"
	"github.com/ebubekir/go-talk/api/pkg/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	env := InitializeEnvironment()

	// Init MongoDB
	// Initialize mongo db
	mongoDb := mongodb.New(env.MongoDbConnectionString, "gotalk")

	// Init Firebase
	firebaseApp, err := firebase.NewFirebaseApp(env.FirebaseProjectId, env.FirebaseConnectionString)
	if err != nil {
		panic(err)
	}

	// Hub
	dispatcher := websocket.NewEventDispatcher()
	roomHub := websocket.NewHub(firebaseApp, dispatcher)

	// User
	userRepo := userInfra.NewMongoDbUserRepository(mongoDb)
	userService := userApp.NewUserService(userRepo)

	// Room
	roomRepo := roomInfra.NewMongoDbRoomRepository(mongoDb)
	roomService := roomApp.NewRoomService(roomRepo, userService)
	roomEventListener := roomApp.NewRoomEventListener(roomService, roomHub)

	// Register event listeners
	dispatcher.Register(roomEventListener)

	// Run hub
	go roomHub.Run()

	// Middlewares
	authMiddleware := middleware.NewAuthMiddleware(firebaseApp, userService)

	// Swagger settings

	switch env.EnvironmentType {
	case Development:
		docs.SwaggerInfo.Title = "Gotalk API [development]"
		docs.SwaggerInfo.Host = "localhost:8080/v1"
		docs.SwaggerInfo.Schemes = []string{"http"}
	case Qa:
		docs.SwaggerInfo.Title = "Gotalk API [qa]"
		docs.SwaggerInfo.Host = "qa.api.gotalk.com/v1"
		docs.SwaggerInfo.Schemes = []string{"https"}
	case Prod:
		docs.SwaggerInfo.Title = "Gotalk API"
		docs.SwaggerInfo.Host = "api.gotalk.com/v1"
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Create api
	api := gin.Default()
	api.Use(CustomRecovery())
	api.Use(Cors())
	api.Use(gin.Logger())
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.GET("/ws/room/:roomId", websocket.RoomWS(roomHub, userService))

	api.Use(authMiddleware.Handler())

	v1Routes := api.Group("/v1")
	{
		v1Routes.Use(authMiddleware.Handler())
		userHttp.RegisterUserRoutes(v1Routes, userService)
		roomHttp.RegisterRoomRoutes(v1Routes, roomService, userService)
	}

	if err := api.Run(":8080"); err != nil {
		panic(err)
	}
}

func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, recovered any) {
		// Handle panic
		msg := "Unhandled Error:"

		if err, hasErr := recovered.(error); hasErr {
			_ = c.Error(err.(error))
			msg = fmt.Sprintf("Unhandled Error: %v", err.(error).Error())
		}
		response.SystemError(c, errors.New(msg))
	})
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	})
}
