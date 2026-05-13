package main

import (
	"bcc-api/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "200 OK")
}

func main() {
	router := gin.Default()

	router.GET("/health", handlers.HealthHandler)
	router.GET("/apps", handlers.GetApps)
	router.POST("/apps", handlers.CreateApp)
	router.Run(":8080")
}
