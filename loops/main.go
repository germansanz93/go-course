package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ { //Forma mas tradicional de usar for
		sum++
	}

	fmt.Println(sum)

	sum = 1

	for sum < 1000 { //Usamos for con una condicion, lo que sucedera es que el bucle se ejecutara hasta que la condicion se deje de cumplir
		sum += 10
		fmt.Println(sum)
	}

	sum = 0
	for { //tambien puedo no poner la condicion y usar un break en determinados casos
		if sum > 1000 {
			break
		}
		sum += 100
		fmt.Println(sum)
	}
}
