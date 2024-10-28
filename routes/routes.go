package routes

import (
	"github.com/labstack/echo/v4"
	"server/controllers"
	"server/middleware"
)

func SetupRoutes(e *echo.Echo) {

	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.Login)

	authRoutes := e.Group("/")

	authRoutes.Use(middleware.AuthMiddleware)

	authRoutes.POST("team", controllers.CreateTeam)
	authRoutes.POST("team/add", middleware.AdminMiddleware(controllers.AddUserToTeam))

	//e.POST("team", middleware.AuthMiddleware(controllers.CreateTeam))
	//e.POST("/team/add", middleware.AuthMiddleware())
	//authRoutes.POST("/team/add", middleware.AdminMiddleware(controllers.AddUserToTeam))
	// authRoutes.POST("/team/remove", middleware.AdminMiddleware(controllers.RemoveUserFromTeam))

}
