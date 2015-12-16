// Copyright © 2014, 2015 Maxim Tishchenko
// All Rights Reserved.

package model

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/qwertmax/timeconverter/cfg"
	"io/ioutil"
	"os"
	"time"
)

// Variable contain SecretKey. SecretKey should be loaded from secret key file from Config.
var SecretKey = "somekey"

// Time to live valid token. After this period - token will be expired.
var TokenLifetime = int64(10)

// Initialize token and time to live.
func InitToken(config cfg.Config) {
	SecretKey, _ = LoadKey(config)
	TokenLifetime = config.TOKEN_LIFETIME
}

// Load token from file. File must be specified in config.
func LoadKey(config cfg.Config) (string, error) {
	if _, err := os.Stat(config.SECRET_KEY); err != nil {
		return "", errors.New("key path not valid")
	}

	key, err := ioutil.ReadFile(config.SECRET_KEY)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

// Generate new token
//
// P.S token is JWT standart.
func TokenGenerate(login, user_id string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["login"] = login
	token.Claims["user_id"] = user_id
	token.Claims["exp"] = time.Now().Add(time.Second * time.Duration(TokenLifetime)).Unix()

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}

// Parse Token from request.
func TokenParse(myToken string) (string, string, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if err == nil && token.Valid {
		return token.Claims["login"].(string), token.Claims["user_id"].(string), nil
	} else {
		return "", "", err
	}
}
