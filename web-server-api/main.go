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

func postList(c *gin.Context) {
	var newList to_do_list

	if err := c.BindJSON(&newList); err != nil {
		return
	}

	tdlists = append(tdlists, newList)
	c.IndentedJSON(http.StatusCreated, newList)
}

func getListByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tdlists {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found!"})
}

func main() {
	router := gin.Default()
	router.GET("/items", getList)
	router.GET("/items/:id", getListByID)
	router.POST("/items", postList)
	router.Run("localhost:8080")
}
