package main

import (
	"fmt"

	"github.com/germansanz93/go-course/interface/vehicles"
)

func main() {
	Display("A")
	Display(12)

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "CAR", "CAR"}
	var d float64
	for _, v := range vArray {
		fmt.Printf("Vehicle %s\n", v)
		veh, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		distance := veh.Distance()
		fmt.Printf("Distance: %.2f\n\n", distance)
		d += distance
	}

	fmt.Printf("Total distance: %.2f\n", d)
}

func Display(value interface{}) { //Al declararlo como interface no importa el tipo de dato, es como si fuera un generic
	fmt.Println(value)
}
