package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type apps struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Size  int    `json:"size"`
}

var app = []apps{
	{ID: "1", Title: "Asphalt", Desc: "Racing", Size: 1900},
	{ID: "2", Title: "Clash Of Clans", Desc: "Startegy", Size: 1300},
	{ID: "3", Title: "Subway Surfers", Desc: "Arcade", Size: 350},
}

func CreateApp(c *gin.Context) {
	var newApp apps

	if err := c.BindJSON(&newApp); err != nil {
		return
	}

	app = append(app, newApp)
	c.IndentedJSON(http.StatusCreated, newApp)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "app recieved"})

}

func GetApps(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, app)
}

func HealthHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
