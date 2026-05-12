package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
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
	id := c.Param("id")

	for _, a := range app {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found!"})
}

func HealthHandler() ([]byte, error) {
	msg := Response{
		Status: "ok",
		Code:   200,
	}
	jsonData, err := json.Marshal(msg)
	return jsonData, err
}

func main() {
	router := gin.Default()
	router.GET("/items", getList)
	router.GET("/items/:id", getListByID)
	router.POST("/items", postList)
	router.Run("localhost:8080")
}
