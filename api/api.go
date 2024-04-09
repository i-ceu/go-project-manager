package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/controllers"
	"github.com/ubaniIsaac/go-project-manager/internal/middleware"
)

//list all api routes from

func RegisterRoutes() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		//auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/signin", controllers.SignIn)
		}

		tasks := v1.Group("/tasks")
		tasks.Use(middleware.AuthMiddleware())
		{
			tasks.POST("/", controllers.CreateTask)
			tasks.GET("/:id", controllers.GetTasks)
		}
		projects := v1.Group("/projects")
		projects.Use(middleware.AuthMiddleware())
		{
			projects.POST("/", controllers.CreateProject)
			projects.GET("/:id", controllers.GetProject)
		}
	}

	r.Run()
}
