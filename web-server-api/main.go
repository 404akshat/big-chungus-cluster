package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type to_do_list struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Time     string `json:"time"`
	Priority int    `json:"priority"`
}

var tdlists = []to_do_list{
	{ID: "1", Title: "Apply for Internships", Time: "10 AM", Priority: 1},
	{ID: "2", Title: "Big Chungus Cluster", Time: "2 PM", Priority: 1},
	{ID: "3", Title: "Coursera", Time: "6 PM", Priority: 2},
}

func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tdlists)
}

func main() {
	router := gin.Default()
	router.GET("/items", getList)
	router.Run("localhost:8080")
}
