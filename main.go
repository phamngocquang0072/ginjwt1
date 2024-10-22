package main

import (

	"github.com/quangnpham/ginjwt1/initializers"
	"github.com/quangnpham/ginjwt1/migrate"
	"github.com/quangnpham/ginjwt1/internal/routers"
	_ "github.com/quangnpham/ginjwt1/docs"
)


func init() {
	
	initializers.LoadEnv()
	initializers.ConnectDB()
	routers.RootRouter()
}
// @title Swagger API
// @description This is a sample server celler server.
// @version 1.0

// @host localhost:5000
// @BasePath /
func main() {
	migrate.Migrate()
}