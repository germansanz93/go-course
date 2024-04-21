package main //declaro el paquete en el que va a estar este archivo

import (
	"fmt" //hago import de fmt para poder usarlo mas adelante
)

func main() { //funcion main, es el punto de entrada de nuestro programa
	var myIntVar int                          //con la palabra var puedo declarar una variable, en este caso de nombre myIntVar y de tipo int
	fmt.Println("My variable is: ", myIntVar) //hago uso del paquete fmt para imprimir una linea por pantalla

	//hasta ahora se imprimio su zero value, pero podria modificarlo
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar) //hago uso del paquete fmt para imprimir una linea por pantalla

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	fmt.Println("My variable is: ", myStringVar)

	myStringVar = "My String Var"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	fmt.Println("My variable is: ", myBoolVar)

	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	fmt.Println("My variable memory address is: ", &myStringVar) //Con & obtengo el puntero hacia una variable

	myIntVar2 := 3 //Si queremos declarar e inicializar de forma rapida una variable lo podemos hacer con la notacion de :=
	fmt.Println("My variable is: ", myIntVar2)

	const myIntConst int = 3

}
