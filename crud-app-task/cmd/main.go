// cmd/main.go
package main

import (
	"crud-app-task/config"
	"crud-app-task/controllers"
	"crud-app-task/repositories"
	"crud-app-task/routes"
	"crud-app-task/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin Engine to handle HTTP
	router := gin.Default()

	// Db connectivity for crud app
	config.ConnectDB()

	// Repository
	userRepository := repositories.NewUserRepository(config.DB)

	//Service
	userService := services.NewUserService(userRepository)

	// Controller
	userController := controllers.NewUserController(userService)

	//Routes
	routes.RegisterUserRoutes(router, userController)

	// Start the Server
	router.Run(":8080")
}
