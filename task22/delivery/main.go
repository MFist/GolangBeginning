package main

import (
	"GolangBeginning/task22/delivery/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/createfile/:filename", handler.CreateFile)
	r.POST("/writefile/:filename", handler.WriteFile)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
