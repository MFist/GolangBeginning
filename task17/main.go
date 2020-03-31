package main

import (
	"GolangBeginning/task17/model"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
)

// Repository
var obj = []model.Game{
	model.Game{Name: "Detroit Become Human", Genre: "Action-adventure", Console: "Playstation 4", Year: 2018}, model.Game{Name: "Horizon Zero Dawn", Genre: "Action role-playing game", Console: "Playstation 4", Year: 2017},
	model.Game{Name: "GTA V", Genre: "Action-adventure", Console: "Multiplataform", Year: 2013},
}

// Business

func randomGame() model.Game {
	game := model.Game{}
	_ = faker.FakeData(&game)
	return game
}

func main() {
	r := gin.Default()
	game := randomGame()

	r.GET("/game", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Game": game,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
