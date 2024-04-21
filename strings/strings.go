package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	var myIntVar int = 3
	myFloatVar := 12.3
	fmt.Printf("Type: %T, value: %d, bytes: %d, bits: %d, \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)
	fmt.Printf("Type: %T, value: %f, bytes: %d, bits: %d, \n", myFloatVar, myFloatVar, unsafe.Sizeof(myFloatVar), unsafe.Sizeof(myFloatVar)*8)

	var myStringVar string = "Hola soy German."
	fmt.Printf("Mi valor es: %s\n", myStringVar)
	fmt.Printf("Type: %T, value: %s, bytes: %d, bits: %d, \n", myStringVar, myStringVar, unsafe.Sizeof(myStringVar), unsafe.Sizeof(myStringVar)*8)

	var myStringVar2 string = `
		Mi variable
		de tipo 
		texto en 
		GO.
	`

	fmt.Printf("Mi valor es: %s\n", myStringVar2)
	fmt.Printf("Type: %T, value: %s, bytes: %d, bits: %d, \n", myStringVar2, myStringVar2, unsafe.Sizeof(myStringVar2), unsafe.Sizeof(myStringVar2)*8)

	//parse to string
	floatVar := 33.11
	fmt.Printf("Type: %T, value: %f\n", floatVar, floatVar)

	floatStrVar := fmt.Sprintf("%.2f", floatVar)

	fmt.Printf("Type: %T, value: %s\n", floatStrVar, floatStrVar)

	intVar2 := 324

	intStrVar2 := strconv.Itoa(intVar2)
	fmt.Printf("Type: %T, value: %s\n", intStrVar2, intStrVar2)

}
