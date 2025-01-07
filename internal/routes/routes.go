package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/i-ceu/go-project-manager/internal/controllers"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/middleware"
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
			auth.GET("/acceptInvite/:id", controllers.AcceptInvite)
			auth.GET("/verify/:id", controllers.VerifyAccount)
		}

		v1.POST("/organization/register", middleware.Guest(), controllers.RegisterOrganization)
		v1.GET("/organization/:organizationId/signin", middleware.Guest(), controllers.SignInToOrganization)

		//organization
		organization := v1.Group("/organization")
		organization.Use(middleware.Auth())
		{
			organization.POST("/invite/:id", controllers.InviteToOrganiztion)
		}

		//projects
		projects := v1.Group("/projects")
		projects.Use(middleware.Auth())
		{
			projects.POST("/", middleware.CheckRole("admin"), controllers.CreateProject)
			projects.GET("all/:organizationId", controllers.GetAllProjects)
			projects.GET("/:id", controllers.GetProject)
		}

		//sprints
		sprint := v1.Group("/sprint")
		sprint.Use(middleware.Auth())
		{
			sprint.POST("/create", controllers.CreateSprint)
		}

		//tasks
		tasks := v1.Group("/tasks")
		tasks.Use(middleware.Auth())
		{
			tasks.POST("/", middleware.CheckRole("admin"), controllers.CreateTask)
			tasks.PUT("/assignTask/:id", controllers.AssignTask)
			tasks.PUT("/updateTask/:id", controllers.UpdateTask)
			tasks.GET("/:id", controllers.GetTask)
		}

	}

	r.Run()
}
