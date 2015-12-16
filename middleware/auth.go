// Copyright © 2014, 2015 Maxim Tishchenko
// All Rights Reserved.

package middleware

import (
	"github.com/autocrm/api/model"
	"github.com/autocrm/api/route"
	"github.com/gin-gonic/gin"
)

// middleware witch require valid token, otherwise microservice return 401 as result.
//
// if token is valid, you will get processed your request.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// token := c.Request.Header.Get("Authenticate")

		if token == "" {
			// c.JSON(401, gin.H{"Status": "erorr", "Message": "Authentication error"})
			// c.Writer.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", realm))
			// c.AbortWithError(401, errors.New("Unauthorized"))
			c.JSON(401, model.Error{
				Type:    "unauthorized",
				Message: "Unauthorized",
			})
			c.AbortWithStatus(401)
			return
		}

		login, user_id, err := model.TokenParse(token)
		if err != nil {
			// c.AbortWithError(401, err)
			c.JSON(401, model.Error{
				Type:    err.Error(),
				Message: err.Error(),
			})
			c.AbortWithStatus(401)
			return
		}

		c.Set("user", login)
		c.Set("user_id", user_id)
		c.Set("isAuth", true)

		c.Next()
	}
}

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["Authorization"]) == 0 {
			// c.AbortWithError(403, errors.New("The API key is invalid."))
			c.JSON(401, model.Error{
				Type:    "the_api_key_is_invalid",
				Message: "The API key is invalid.",
			})
			c.AbortWithStatus(401)
			return
		}

		key, err := route.ParseAuthKey(c)
		if err != nil {
			// c.AbortWithError(403, errors.New("Error Parse key."))
			c.JSON(401, model.Error{
				Type:    "error_parse_key",
				Message: "Error Parse key",
			})
			c.AbortWithStatus(401)
			return
		}

		Db := route.GetDB(c)
		apikey := model.ApiKey{}

		var num int
		Db.Where(&model.ApiKey{
			Name: key,
		}).First(&apikey).Count(&num)

		if num == 0 {
			// c.AbortWithError(403, errors.New("The API key is not found."))
			c.JSON(401, model.Error{
				Type:    "the_api_key_is_not_found",
				Message: "The API key is not found",
			})
			c.AbortWithStatus(401)
			return
		}

		if apikey.HasBegunWorking() {
			c.JSON(401, model.Error{
				Type:    "api_key_havent_begun_working",
				Message: "API key haven't begun working",
			})
			c.AbortWithStatus(401)
			return
		}

		if apikey.HasExpired() {
			c.JSON(401, model.Error{
				Type:    "api_key_has_expired",
				Message: "API Key has expired",
			})
			c.AbortWithStatus(401)
			return
		}

		c.Set("key_id", apikey.ID)
		c.Next()
	}
}
