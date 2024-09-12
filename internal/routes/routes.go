package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/controllers"
	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
	"github.com/ubaniIsaac/go-project-manager/internal/middleware"
)

//list all api routes from

func RegisterRoutes() {

	logFile := helpers.SetupLogging()
	defer logFile.Close()

	gin.DefaultWriter = logFile

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		//auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/registerUser", controllers.RegisterUser)
			auth.POST("/signin", controllers.SignIn)
			auth.POST("/acceptInvite/:id", controllers.AcceptInvite)
		}

		//organization
		organization := v1.Group("/organization")
		organization.Use(middleware.Auth())
		{
			organization.POST("/invite/:id", controllers.InviteToOrganiztion)
			organization.POST("/register", middleware.CheckRole("admin"), controllers.RegisterOrganization)
		}

		tasks := v1.Group("/tasks")
		tasks.Use(middleware.Auth())
		{
			tasks.POST("/", middleware.CheckRole("PM"), controllers.CreateTask)
			tasks.PUT("/assignTask/:id", controllers.AssignTask)
			tasks.PUT("/updateTask/:id", controllers.UpdateTask)
			tasks.GET("/:id", controllers.GetTask)
		}

		projects := v1.Group("/projects")
		projects.Use(middleware.Auth())
		{
			projects.GET("/", controllers.GetAllProjects)
			projects.POST("/", middleware.CheckRole("PM"), controllers.CreateProject)
			projects.GET("/:id", controllers.GetProject)
		}
	}

	r.Run()
}
