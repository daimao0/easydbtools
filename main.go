package main

import (
	"easydbTools/internal/adapter/http"
	"github.com/gin-gonic/gin"
)

// This is a DB GUI written in Go + gin + gorm

func main() {
	engine := gin.Default()
	http.RegisterRoutes(engine)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
