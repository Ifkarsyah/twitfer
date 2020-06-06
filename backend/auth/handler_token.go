package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func Refresh(c *gin.Context) {
	refreshToken, done := getOldRefreshToken(c)
	if done {
		return
	}
	token, err := jwt.Parse(refreshToken, checkConformHMAC("REFRESH_TOKEN"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, "the token claims should conform to MapClaims")
		return
	}

	//Since token is valid, get the uuid:
	refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Error occurred")
		return
	}

	//Delete the previous Refresh Token
	deleted, delErr := RedisDeleteAuth(refreshUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	//Create new pairs of refresh and access tokens
	ts, createErr := CreateToken(userId)
	if createErr != nil {
		c.JSON(http.StatusForbidden, createErr.Error())
		return
	}

	//save the tokens metadata to redis
	saveErr := RedisCreateAuth(userId, ts)
	if saveErr != nil {
		c.JSON(http.StatusForbidden, saveErr.Error())
		return
	}


	c.JSON(http.StatusCreated, map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	})
}

func getOldRefreshToken(c *gin.Context) (string, bool) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return "", true
	}
	return mapToken["refresh_token"], false
}
