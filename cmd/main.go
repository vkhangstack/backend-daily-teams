package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vkhangstack/dlt/internal/adapters/handler"
	"github.com/vkhangstack/dlt/internal/adapters/middlewares"
	"github.com/vkhangstack/dlt/internal/adapters/repository"
	"github.com/vkhangstack/dlt/internal/adapters/utils"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"github.com/vkhangstack/dlt/internal/core/services"
	"github.com/vkhangstack/dlt/internal/logger"
	"os"
)

var (
	userService  *services.UserService
	dailyService *services.DailyService
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	gin.SetMode(os.Getenv("MODE"))
	gin.Recovery()

	db, err := repository.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	//redisCache, err := cache.NewRedisCache("127.0.0.1:6379", "")
	//if err != nil {
	//	panic(err)
	//}

	// Create or modify the database tables based on the model structs found in the imported package
	db.AutoMigrate(&model.User{}, &model.DailyTask{})

	//store := repository.NewDB(db, redisCache)
	store := repository.NewDB(db)
	//firebasePath := os.Getenv("FIREBASE_PATH")

	//newFirebaseApp, _ := firebaseApp.NewFirebaseApp(firebasePath)

	//msgService = services.NewMessengerService(store)
	userService = services.NewUserService(store)
	dailyService = services.NewDailyService(store)

	utils.NewSnowflakeService(1)

	InitRoutes()
}

func InitRoutes() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	v1 := router.Group("/v1")

	userGroup := v1.Group("/users")
	userGroup.Use(middlewares.AuthMiddleware(*userService))

	userHandler := handler.NewUserHandler(*userService)

	userGroup.Use(middlewares.AuthMiddleware(*userService))
	userGroup.GET("/profile/me", userHandler.ProfileMe)

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", userHandler.Register)
	authGroup.POST("/login", userHandler.LoginWithKey)
	authGroup.GET("/access", userHandler.GetAccessToken)

	dailyGroup := v1.Group("/daily")
	dailyGroup.Use(middlewares.AuthMiddleware(*userService))

	dailyHandler := handler.NewDailyHandler(*dailyService)

	dailyGroup.POST("/task", dailyHandler.CreateTask)
	dailyGroup.PUT("/task", dailyHandler.UpdateTask)
	dailyGroup.DELETE("/task/:id", dailyHandler.DeleteTask)
	dailyGroup.GET("/task", dailyHandler.ListTasks)

	logger.Log.Infoln("Server started listening on :4000")
	err := router.Run(":4000")

	if err != nil {
		logger.Log.Errorf("Error starting server: %v", err)
	}

	// go func() {
	// 	if err := router.Run(":5000"); err != nil {
	// 		log.Fatalf("failed to run messages and users service: %v", err)
	// 	}
	// }()

}
