package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qwertmax/timeconverter/cfg"
	// "github.com/qwertmax/timeconverter/db"
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

	//Middleware
	r.Use(db.DB(&database))

	// r.LoadHTMLGlob("templates/*")
	r.HTMLRender = CreateMyRender()
	r.Static("/assets", "./assets")

	//routings
	r.GET("/", route.Main)
	r.GET("/uploadTeams", route.UploadTeamsHTML)
	r.GET("/uploadEvents", route.UploadEventsHTML)
	r.GET("/uploadPlayers", route.UploadPlayersHTML)

	r.GET("/event/html/:id", route.EventHTML)
	r.GET("/team/add", route.TeamAddHTML)
	r.GET("/team", route.TeamsHTML)
	r.GET("/event", route.EventsHTML)
	// r.GET("/event/add", route.EventAddHTML)
	r.GET("/event/edit/:id", route.EventEditHTML)
	// r.GET("/event/:id", route.EventHTML)

	api := r.Group("/api")
	{
		api.POST("/uploadTeams", route.UploadTeams)
		api.POST("/uploadEvents", route.UploadEvents)
		api.POST("/uploadPlayers", route.UploadPlayers)

		api.GET("/events", route.EventList)
		api.GET("/event/:id", route.EventGet)
		api.POST("/event/:id", route.EventCreate)
		api.PUT("/event/:id", route.EventUpdate)
		api.GET("/events/count", route.EventsCount)
		api.GET("/events/count/:year", route.EventsCountYear)

		api.POST("/team", route.TeamAdd)
		api.DELETE("/team/:id", route.TeamDelete)

		api.GET("/states", route.GetStates)

		api.GET("/sport", route.Sport)
		api.GET("/sport-fill", route.SportFill)
	}

	// r.POST("/api/uploadTeams", route.UploadTeams)
	// r.POST("/api/uploadEvents", route.UploadEvents)
	// r.POST("/api/uploadPlayers", route.UploadPlayers)

	// r.POST("/api/team", route.TeamAdd)
	// r.DELETE("/api/team/:id", route.TeamDelete)

	// r.GET("/api/states", route.GetStates)
	// r.GET("/api/sport", route.Sport)
	// r.GET("/api/sport-fill", route.SportFill)

	// r.GET("/api/events", route.EventList)
	// r.GET("/api/event/:id", route.EventGet)
	// r.POST("/api/event/:id", route.EventCreate)
	// r.PUT("/api/event/:id", route.EventUpdate)
	// r.GET("/api/events/count", route.EventsCount)
	// r.GET("/api/events/count/:year", route.EventsCountYear)

	r.GET("/v1/team", route.TeamsAPI)
	r.GET("/v1/player", route.Playres)

	// r.GET("/team/:id/powerrank", route.TeamPowerrank)

	// start service
	r.Run(":" + config.APP_PORT)
}
