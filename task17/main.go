package main

import (
	"GolangBeginning/task17/model"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	game := model.Game{}
	err := faker.FakeData(&game)

	r.GET("/game", func(c *gin.Context) {
		if err != nil {
			c.JSON(500, gin.H{
				"Message": "Bad request",
			})
		}
		c.JSON(200, gin.H{
			"Game": game,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
