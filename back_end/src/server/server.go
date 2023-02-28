package server

import (
	"CEN3031-Project/back_end/src/controllers"
	"CEN3031-Project/back_end/src/middlewares"

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
		AllowOrigins: []string{"http://localhost:4200", "http://localhost:1337"},
		// AllowHeaders:     []string{"Origin"},
		// ExposeHeaders:    []string{"Content-Length"},
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
			secured.DELETE("/token", controllers.DeleteUser)
			secured.POST("/createCourse", controllers.CreateCourse)
			secured.POST("/addAssignment", controllers.CreateAssignment)
			secured.POST("/completeAssignment", controllers.CompleteAssignment)
		}
		users.GET("/", controllers.GetUsers)
	}
	//Achievment group
	achievements := r.Group("/achievements")
	{
		achievements.POST("/createAchievement", controllers.AddAchievement)
		achievements.GET("/getAllAchievements", controllers.GetAllAchievements)
		//achievements.DELETE("/deleteAchievement", controllers.DeleteAchievement)
		//achievements.PUT("/editAchievement", controllers.EditAchievement)
	}
	return r
}
