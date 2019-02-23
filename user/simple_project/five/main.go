package main

import "fmt"

type car struct {
	gasPedal      uint16
	breakPedal    uint16
	steeringWheel int16
	topSpeed      float64
}

func main() {
	bmw := car{
		gasPedal:      10,
		breakPedal:    25,
		steeringWheel: 100,
		topSpeed:      250,
	}
	fmt.Printf("THis is struct %v", bmw)
}
