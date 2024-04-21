package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d\n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar1 [5]int
	fmt.Println(myArrayVar1, "- size:", len(myArrayVar1))

	myArrayVar2 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar2, "- size:", len(myArrayVar2))

	myArrayVar1[0] = 5
	myArrayVar1[1] = 4
	myArrayVar1[2] = 3
	myArrayVar1[3] = 2
	myArrayVar1[4] = 1
	fmt.Println(myArrayVar1, "- size:", len(myArrayVar1))
	fmt.Println(myArrayVar1[1])
	fmt.Printf("type: %T, value: %v, bytes: %d, bits: %d\n", myArrayVar1, myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myIntVar)*8)

	for i := range myArrayVar1 {
		fmt.Println("index:", i, "- value:", myArrayVar1[i])
	}

	for i, v := range myArrayVar1 {
		fmt.Println("index:", i, "- value:", v)
	}

	for _, v := range myArrayVar1 {
		fmt.Println("value:", v)
	}
}
