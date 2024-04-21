package main

import "fmt"

func main() {
	yearsOld := 33

	if yearsOld > 18 {
		fmt.Printf("%d es mayor a 18\n", yearsOld)
	}

	boolVar := true

	if boolVar {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es falso")
	}

	if value := true; value { //Puedo dentro de un if declarar una variable, luego usar punto y coma y por ultimo poner el valor para que sea validado por el if
		// tal vez en el ejemplo no tiene tanto sentido pero es muy util cuando tenemos que evaluar
		// nuestra condicion en base a por ejemplo una funcion.
		fmt.Println("es verdadero el if raro")
	}

	number := 3

	if number == 1 {
		fmt.Println("es 1")
	} else if number == 2 {
		fmt.Println("es 2")
	} else if number == 3 {
		fmt.Println("es 3")
	} else {
		fmt.Println("ninguna es valida")
	}

	switch number {
	case 1:
		fmt.Println("es 1")
	case 2:
		fmt.Println("es 2")
	case 3:
		fmt.Println("es 3")
	default:
		fmt.Println("ninguna es valida")
	}

	switch number2 := 5; number2 { //Con el switch tambien puedo asignar un valor y validarlo
	case 1:
		fmt.Println("es 1")
	case 2:
		fmt.Println("es 2")
	case 3:
		fmt.Println("es 3")
	default:
		fmt.Println("ninguna es valida")
	}

	switch { //Otra forma de usar el switch es evaluando directamente la condicion en el case
	case number > 0:
		fmt.Println("es positivo")
	case number == 0:
		fmt.Println("es 0")
	default:
		fmt.Println("es negativo")
	}

}
