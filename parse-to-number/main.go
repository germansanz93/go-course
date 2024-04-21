package main

import (
	"fmt"
	"strconv"
)

func main() {
	strIntVar, err := strconv.Atoi("15")
	fmt.Printf("type: %T, value: %d, err: %v\n", strIntVar, strIntVar, err)

	strIntVar2, err := strconv.Atoi("15a")
	fmt.Printf("type: %T, value: %d, err: %v\n", strIntVar2, strIntVar2, err)

	strIntVar3, _ := strconv.ParseInt("111", 2, 64) //La funcion ParseInt del modulo strconv es algo diferente a Atoi, ya que nos permite elegir la base en la que esta el entero en el string y su tamano. En este caso base 10 y tamano 64bits
	fmt.Printf("type: %T, value: %d\n", strIntVar3, strIntVar3)

	strFloatVar, err := strconv.ParseFloat("-11.2", 64) //Indico el tamano como segundo arg
	fmt.Printf("type: %T, value: %f\n", strFloatVar, strFloatVar)

}
