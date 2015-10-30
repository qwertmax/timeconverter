package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qwertmax/timeconverter/model"
)

func Main(c *gin.Context) {
	Db := GetDB(c)

	var evens []model.Event
	err := Db.Limit(10).Find(&evens).Order("date desc").Error
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.HTML(200, "index", gin.H{
		"content": "upload_event.tmpl",
		"events":  evens,
	})
}

func GetStates(c *gin.Context) {
	Db := GetDB(c)

	type State struct {
		State string `json:"State" sql:"state"`
		Abr   string `json:"State_Abr" sql:"state_abr"`
	}
	var states []State

	rows, err := Db.Table("team").Select("distinct state, state_abr").Rows()
	if err != nil {
		c.AbortWithError(500, err)
	}

	for rows.Next() {
		var state State
		rows.Scan(&state.State, &state.Abr)
		states = append(states, state)
	}

	c.JSON(200, states)
}
