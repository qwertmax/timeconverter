package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qwertmax/timeconverter/cfg"
	"github.com/qwertmax/timeconverter/db"
	"github.com/qwertmax/timeconverter/middleware"
	"github.com/qwertmax/timeconverter/route"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	// gin.SetMode(gin.ReleaseMode)

	// model.InitToken(config)
	// config := cfg.Init()
	config, err := cfg.GetConfig("conf.yml")
	if err != nil {
		panic(err)
	}

	if config.RELEASE_MODE {
		gin.SetMode(gin.ReleaseMode)
	}
	fmt.Println(config)

	var database db.Database
	database.Init(config)

	// create routing
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	//Middleware
	r.Use(db.DB(&database))

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/users", route.UsersList)
		authorized.GET("/user/:id", route.UserGet)
		authorized.POST("/user", route.UserCreate)
		authorized.PUT("/user/:id", route.UserUpdate)
		authorized.DELETE("/user/:id", route.UserDelete)

	}

	//routings
	r.GET("/", route.Main)

	// r.GET("/users", route.UsersList)
	// r.GET("/user/:id", route.UserGet)
	// r.POST("/user", route.UserCreate)
	// r.PUT("/user/:id", route.UserUpdate)
	// r.DELETE("/user/:id", route.UserDelete)

	// start service
	r.Run(":" + config.APP_PORT)
}
