package server

import (
	"AttackOnCollege/back_end/src/controllers"
	"AttackOnCollege/back_end/src/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := initRouter()
	r.Run(":1337")

}

func initRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:1337", "http://localhost:9876/"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
	users := r.Group("/users")
	{
		users.POST("/token", controllers.GenerateToken)
		users.POST("/register", controllers.RegisterUser)
		secured := users.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/token", controllers.GetUser)
			secured.GET("/achievements", controllers.GetAchievements)
			secured.PUT("/token", controllers.EditUser)
			secured.POST("/logout", controllers.LogOut)
			secured.PUT("/assignment", controllers.EditAssignment)
			secured.DELETE("/token", controllers.DeleteUser)
			secured.POST("/course", controllers.CreateCourse)
      secured.PUT("/course", controllers.EditCourse)
      secured.POST("/coursedone", controllers.CompleteCourse)
			secured.POST("/assignment", controllers.CreateAssignment)
			secured.POST("/complete", controllers.CompleteAssignment)
			secured.GET("/assignments", controllers.GetAssignments)
			secured.GET("/assignment", controllers.GetAssignment)
		}
		admin := users.Group("/admin").Use(middlewares.Auth())
		{
			admin.POST("/achievement", controllers.AddAchievement)
			admin.PUT("/achievement", controllers.EditAchievement)
			admin.DELETE("/achievement", controllers.DeleteAchievement)
			admin.GET("/achievements", controllers.GetAllAchievements)
			//admin.GET("/achievement", controllers.ReturnAchievement)

			//admin.DELETE("/user", controllers.AdminDeleteUser)
			//admin.GET("/users", controllers.AdminGetAllUsers)
			//admin.PUT("/user", controllers.AdminEditUser)
		}
		users.GET("/", controllers.GetUsers)
	}

	return r
}
