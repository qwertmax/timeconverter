// Copyright © 2015, 2016 Maxim Tishchenko
// All Rights Reserved.

// Package cfg implements a simple config parser from YML config file.
package cfg

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// Structure for yaml config parameters:
// 		APP_PORT			- application port
//		DB_USERNAME			- username for DB connection
//		DB_PASSWORD			- paassword for DB connection
//		DB_NAME				- Database Name for DB connection
//		DB_ADDRESS			- IP address for DB connection
//		DB_SSLMODE			- enable/disable SSH mode in DB connection
//		DB_PORT				- port for DB connection
// 		SECRET_KEY 			- path to secret kry
// 		TOKEN_LIFETIME 		- time in sedonds for vailid token.
//		ACCESS_CONTROL_ALLOW_ORIGIN	- this option for CORS requests see https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS
type Config struct {
	APP_PORT     string `yaml:"app_port"`
	DB_USERNAME  string `yaml:"db_username"`
	DB_PASSWORD  string `yaml:"db_password"`
	DB_NAME      string `yaml:"db_name"`
	DB_ADDRESS   string `yaml:"db_address"`
	DB_SSLMODE   string `yaml:"db_ssh_mode"`
	DB_PORT      string `yaml:"db_port"`
	RELEASE_MODE bool   `yaml:"release_mode"`
	DB_LOG       bool   `yaml:"db_log"`
}

// initialize confi, load and perse it, or user defaukr values.
func Init() Config {
	filename := "conf.yml"

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		file, err = os.Create(filename)
		CheckErr(err, "os.Create:")
		defer file.Close()
		config := Config{
			APP_PORT:     "3000",
			DB_USERNAME:  "oddsmaker",
			DB_PASSWORD:  "1",
			DB_NAME:      "oddsmaker",
			DB_ADDRESS:   "192.168.99.100",
			DB_SSLMODE:   "SSLMode",
			DB_PORT:      "5432",
			RELEASE_MODE: false,
			DB_LOG:       false,
		}

		data, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		file.Write(data)

		return config
	}

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		CheckErr(err, "error read config!: ")
	}
	return config
}

func GetConfig(yamlPath string) (Config, error) {
	// yamlPath := c.GlobalString("config")
	config := Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(ymlData, &config)
	if err != nil {
		return config, err
	}

	dbAddress := os.Getenv("DB_PORT_5432_TCP_ADDR")
	if len(dbAddress) > 0 {
		config.DB_ADDRESS = dbAddress
	}

	dbPort := os.Getenv("DB_PORT_5432_TCP_PORT")
	if len(dbPort) > 0 {
		config.DB_PORT = dbPort
	}

	dbUser := os.Getenv("DB_ENV_POSTGRES_USER")
	if len(dbUser) > 0 {
		config.DB_USERNAME = dbUser
	}

	dbPassword := os.Getenv("DB_ENV_POSTGRES_PASSWORD")
	if len(dbPassword) > 0 {
		config.DB_PASSWORD = dbPassword
	}

	// err = config.Test()
	// if err != nil {
	// 	return config, err
	// }

	return config, nil
}

// TODO: add more tests
func (c *Config) Test() error {
	f, err := os.Open(c.SECRET_KEY)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}

	fileLenght := file.Size()
	if fileLenght == 0 {
		return errors.New("secret_key is 0 bytes")
	}

	if fileLenght < 20 {
		return errors.New("secret_key is less than 20 bytes")
	}

	return nil
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
