package server

import (
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/controllers"
	"CEN3031-Project/AttackOnCollege_v0.0.1/src/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := initRouter()

	httpRouter := gin.Default()

	httpRouter.GET("/:path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:443/"+c.Param("path"))
	})
	httpRouter.POST("/*path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:443/"+c.Param("path"))
	})
	httpRouter.PUT("/:path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:443/"+c.Param("path"))
	})
	httpRouter.DELETE("/:path", func(c *gin.Context) {
		c.Redirect(302, "https://localhost:443/"+c.Param("path"))
	})

	go r.RunTLS(":1337", "./server/auth/cert/cacert.crt", "./server/auth/cert/ca.key")
	httpRouter.Run(":8080")

}

func initRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	users := r.Group("/users")
	{
		users.POST("/token", controllers.GenerateToken)
		users.POST("/user/register", controllers.RegisterUser)
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
	return r
}
