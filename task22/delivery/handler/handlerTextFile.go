package handler

import (
	"GolangBeginning/task22/repository"

	"github.com/gin-gonic/gin"
)

// CreateFile this is the handler of the CreateFile function
func CreateFile(c *gin.Context) {
	filename := c.Param("filename")

	createFile := repository.CreateFile(filename)

	c.JSON(200, gin.H{
		"Message": createFile,
	})
}

// WriteFile this is the handler of the CreateFile function
func WriteFile(c *gin.Context) {
	filename := c.Param("filename")

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	writeFile := repository.WriteFile(filename, reqBody)

	c.JSON(200, gin.H{
		"Message": writeFile,
	})
}
