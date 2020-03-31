package main

import (
	"GolangBeginning/task17/model"
	"fmt"

	"github.com/bxcodec/faker"
)

func main() {
	a := model.Game{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
