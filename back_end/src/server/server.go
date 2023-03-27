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
			secured.PUT("/token", controllers.EditUser)
			secured.PUT("/editAssignment", controllers.EditAssignment)
			secured.DELETE("/token", controllers.DeleteUser)
			secured.POST("/createCourse", controllers.CreateCourse)
			secured.POST("/addAssignment", controllers.CreateAssignment)
			secured.POST("/completeAssignment", controllers.CompleteAssignment)
		}
		admin := users.Group("/admin").Use(middlewares.Auth())
		{
			admin.POST("/achievement", controllers.AddAchievement)
			admin.PUT("/achievement", controllers.EditAchievement)
			admin.DELETE("/achievement", controllers.DeleteAchievement)
			admin.GET("/achievements", controllers.GetAllAchievements)
			admin.DELETE("/user", controllers.AdminDeleteUser)
		}
		users.GET("/", controllers.GetUsers)
	}
	//Achievment group
	achievements := r.Group("/achievements")
	{
		achievements.POST("/createAchievement", controllers.AddAchievement)
		achievements.GET("/getAllAchievements", controllers.GetAllAchievements)
		achievements.DELETE("/deleteAchievement", controllers.DeleteAchievement)
		achievements.PUT("/editAchievement", controllers.EditAchievement)
	}
	return r
}
