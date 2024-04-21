package main

import "fmt"

func main() {
	myArrayVar := [5]int{3, 6, 9, 10, 16}
	fmt.Printf("Type: %T", myArrayVar)
	fmt.Println("Array:", myArrayVar, "- len:", len(myArrayVar))

	mySliceVar := []int{}

	fmt.Printf("Type: %T\n", mySliceVar)
	fmt.Println("Slice:", mySliceVar, "- len:", len(mySliceVar))

	mySliceVar = append(mySliceVar, 12, 34, 54)
	fmt.Println("Slice:", mySliceVar, "- len:", len(mySliceVar))

	mySliceVar2 := myArrayVar[2:4]
	fmt.Println("Slice:", mySliceVar2, "- len:", len(mySliceVar2), "- address: ", &mySliceVar2[1])

	mySliceVar2[1] = 19 //OJO!!
	fmt.Println("Array:", myArrayVar, "- len:", len(myArrayVar), "- address: ", &myArrayVar[3])

	mySliceVar3 := myArrayVar[:4] //esta notacion arranca desde el cero
	fmt.Println("Array:", mySliceVar3, "- len:", len(mySliceVar3))

	mySliceVar4 := myArrayVar[2:] //esta notacion indica hasta el final
	fmt.Println("Array:", mySliceVar4, "- len:", len(mySliceVar4))

	//Tambien podemos inicializar un slice con la palabra reservada make, indicando el tipo y el tamano
	mySliceVar5 := make([]int, 5)
	fmt.Println("Slice:", mySliceVar5, "- len:", len(mySliceVar5))
	fmt.Printf("Type: %T\n", mySliceVar5)

	mySliceVar6 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", mySliceVar6, "- len:", len(mySliceVar6))
	fmt.Printf("Type: %T\n", mySliceVar6)
}
