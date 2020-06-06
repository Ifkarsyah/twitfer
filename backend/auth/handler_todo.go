package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

func CreateTodo(c *gin.Context) {
	var td *Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, ErrorResp{List: err.Error()})
		return
	}

	// Who make this request?
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userId, err := RedisGetAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	// Business Logic: I will do this for you, but not know
	td.UserID = userId
	if err := Publish("add_q", c.GetRawData()); err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, td)
}

