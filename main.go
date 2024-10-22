package main

import (

	"github.com/phamngocquang0072/ginjwt1/migrate"
	"github.com/phamngocquang0072/ginjwt1/initializers"
	"github.com/phamngocquang0072/ginjwt1/internal/routers"
	_ "github.com/phamngocquang0072/ginjwt1/docs"
)


func init() {
	
	initializers.LoadEnv()
	initializers.ConnectDB()
	migrate.Migrate()
	routers.RootRouter()
}
// @title Swagger API
// @description This is a sample server celler server.
// @version 1.0

// @host localhost:5000
// @BasePath /
func main() {
}