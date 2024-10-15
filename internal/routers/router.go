package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/phamngocquang0072/gintwj1/internal/controllers"
)

func RootRouter(){
	port := os.Getenv("PORT")
	r := gin.Default()
	// UserRouter(r)
	r.Run(":" + port)	
}

// func UserRouter(r *gin.Engine){
// 	r.GET("/user", controllers.GetUser)
// 	r.POST("/user", controllers.CreateUser)
// }	
	