package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.LoadHTMLGlob("assets/*.html")
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	http.Handle("/", r)

	log.Println("Starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
