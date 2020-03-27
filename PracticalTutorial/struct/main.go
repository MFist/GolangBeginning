package main

import "fmt"

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKm    float64
}

func main() {
	firstCar := car{gasPedal: 22341, brakePedal: 0, steeringWheel: 12561, topSpeedKm: 225.0}

	fmt.Println(firstCar.gasPedal)
}
