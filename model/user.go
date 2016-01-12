package model

type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Cities   []City `json:"cities"`
}

type City struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	TimeZone int64  `json:"time_zone"`
}
