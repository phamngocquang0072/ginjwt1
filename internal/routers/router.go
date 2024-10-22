package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/phamngocquang0072/ginjwt1/internal/controllers"
	"github.com/phamngocquang0072/ginjwt1/internal/middlewares"
    swaggerFiles "github.com/swaggo/files" // swagger embed files
    ginSwagger "github.com/swaggo/gin-swagger" // swagger gin middleware
)

func RootRouter(){
	port := os.Getenv("PORT")
	r := gin.Default()
	UserRouter(r)
	r.Run(":" + port)	
}

func UserRouter(r *gin.Engine){
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.DELETE("/user/:id", controllers.DeleteUser)
}	
	