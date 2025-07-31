package routes

import (
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")
	{
		//auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/registerUser", controllers.RegisterUser)
			auth.POST("/signin", controllers.SignIn)
			auth.POST("/acceptInvite/:inviteId", controllers.AcceptInvite)
			auth.GET("/verify/:verificationId", controllers.VerifyAccount)
		}

		v1.POST("/team/register", middleware.Guest(), controllers.RegisterTeam)
		v1.GET("/team/allTeams", middleware.Guest(), controllers.GetAllUserTeams)
		v1.GET("/team/:teamId/signin", middleware.Guest(), controllers.SignInToTeam)
		v1.GET("/roles", controllers.GetRoles)

		//team
		team := v1.Group("/team")
		team.Use(middleware.Auth())
		{
			team.POST("/invite/:teamId", middleware.CheckRole("admin"), controllers.InviteToTeam)
			team.GET("/:teamId/members", controllers.GetTeamMembers)

			//projects
			projects := team.Group("/:teamId/project")
			{
				projects.POST("/", middleware.CheckRole("admin"), controllers.CreateProject)
				projects.GET("/", controllers.GetAllProjects)
				projects.GET("/:projectId", controllers.GetProject)

				//sprints
				sprint := projects.Group("/:projectId/sprint")
				{
					sprint.POST("/", middleware.CheckRole("admin"), controllers.CreateSprint)
					// sprint.GET("/", controllers.GetAllSprints)
					// sprint.GET("/:sprintId", controllers.GetSprint)
					// sprint.PUT("/:sprintId", middleware.CheckRole("admin"), controllers.UpdateSprint)
					// sprint.DELETE("/:sprintId", middleware.CheckRole("admin"), controllers.DeleteSprint)
					// sprint.POST("/:sprintId/start", middleware.CheckRole("admin"), controllers.StartSprint)
					// sprint.POST("/:sprintId/complete", middleware.CheckRole("admin"), controllers.CompleteSprint)

				}
				//tasks
				tasks := projects.Group("/:projectId/task")
				{
					tasks.POST("/", middleware.CheckRole("admin"), controllers.CreateTask)
					tasks.PUT("/assignTask/:taskId", controllers.AssignTask)
					tasks.PUT("/updateTask/:taskId", controllers.UpdateTask)
					tasks.GET("/:id", controllers.GetTask)
					tasks.GET("/", controllers.GetAllTasks)
				}
			}
		}

	}

	r.Run()
}
