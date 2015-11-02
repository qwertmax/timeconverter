// Copyright © 2014, 2015 Maxim Tishchenko
// All Rights Reserved.

// Package db implements a DB connection and DB operations
//
// It user config parameters:
//
//		DB_USERNAME			- username for DB connection
//		DB_PASSWORD			- paassword for DB connection
//		DB_NAME				- Database Name for DB connection
//		DB_ADDRESS			- IP address for DB connection
//		DB_SSLMODE			- enable/disable SSH mode in DB connection
//		DB_PORT				- port for DB connection
package db

import (
	"fmt"
	_ "github.com/bmizerany/pq"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qwertmax/timeconverter/cfg"
	"github.com/qwertmax/timeconverter/model"
	"time"
)

// Global Dababase Variable for permanent connection.

type Database struct {
	DB *gorm.DB
}

func (db *Database) getDb(conf cfg.Config) (gorm.DB, error) {
	dbconn := "user=" + conf.DB_USERNAME + " password=" + conf.DB_PASSWORD + " dbname=" + conf.DB_NAME + " sslmode=" + conf.DB_SSLMODE + " host=" + conf.DB_ADDRESS + " port=" + conf.DB_PORT
	return gorm.Open("postgres", dbconn)
}

// Initialize DB
func (db *Database) Init(conf cfg.Config) {
	// dbHandler, err := db.getDb(conf)

	// if err != nil {
	// 	panic(err)
	// }

	timeout := make(chan bool, 1)
	defer close(timeout)

	go func() {
		i := 0
		for {
			i++
			fmt.Println("trying connect to database, try #", i)
			dbHandler, _ := db.getDb(conf)
			err := dbHandler.DB().Ping()
			if err == nil {
				timeout <- true
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-timeout:
		dbHandler, err := db.getDb(conf)
		if err != nil {
			panic(err)
		}

		db.DB = &dbHandler

		dbHandler.DB()
		err = dbHandler.DB().Ping()
		if err != nil {
			panic(err)
		}
		dbHandler.DB().SetMaxIdleConns(10)
		dbHandler.DB().SetMaxOpenConns(100)

		// Disable table name's pluralization
		dbHandler.SingularTable(true)

		if conf.DB_LOG {
			dbHandler.LogMode(true)
		}

		dbHandler.CreateTable(&model.User{})
		dbHandler.CreateTable(&model.City{})
		dbHandler.Model(&model.User{}).AddUniqueIndex("idx_email", "email")

		return
	}

}

// Connect to DB and return DB connection handler
func DB(db *Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
