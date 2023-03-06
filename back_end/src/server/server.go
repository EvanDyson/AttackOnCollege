package server

import (
<<<<<<< HEAD
	"CEN3031-Project/back_end/src/controllers"
	"CEN3031-Project/back_end/src/middlewares"
	"os"
=======
	"AttackOnCollege/back_end/src/controllers"
	"AttackOnCollege/back_end/src/middlewares"
>>>>>>> cb9a4465067009ba0ee2e0bc97ff78efc64885e7

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := initRouter()
<<<<<<< HEAD

	httpRouter := gin.Default()

	httpRouter.GET("/:path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:1337/"+c.Param("path"))
	})
	httpRouter.POST("/*path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:1337/"+c.Param("path"))
	})
	httpRouter.PUT("/:path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:1337/"+c.Param("path"))
	})
	httpRouter.DELETE("/:path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:1337/"+c.Param("path"))
	})

	// Change paths once we start running the server with npm
	go r.RunTLS(":"+os.Getenv("PORT"), "./back_end/src/server/auth/cert/cacert.crt", "./back_end/src/server/auth/cert/ca.key")
	httpRouter.Run(":8080")
=======
	r.Run(":1337")
>>>>>>> ce53c2e2a020a8a5009001757e0c168bbe19f4e0

}

func initRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:1337", "http://localhost:9876/"},
		AllowHeaders:     []string{"Origin"},
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
