package function

import (
	"errors"
	"fmt"
)

type Operation int //Defino un tipo Operation que va a ser representado por un int

const ( //Defino constantes del tipo Operation y al hacerlo de la siguiente forma, lo que estoy haciendo en realidad es definir un enum
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Display(myValue int64) {
	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v\n", myValue, myValue, &myValue)
	fmt.Println()
}

func Add(x int, y int) int { //Funcion con valor de retorno de tipo int
	t := x + y
	return t
}

func Add2(x, y int) int { //Cuando tenemos varios valores del mismo tipo podemos solo indicar el tipo en el ultimo
	return x + y
}

func RepeatString(times int, value string) {
	for i := 0; i < times; i++ {
		fmt.Print(value)
	}
	fmt.Println()
}

func Calc(op Operation, x float64, y float64) (float64, error) { //Go nos permite retornas mas de un valor
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("ZeroDivisionError") //Generamos un error nuestro usando el package errors
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}

	return 0, errors.New("InvalidOperationError")
}

func Split(v int) (x, y int) { //Podemos especificar el nombre de los valores a retornar y no hace falta indicar los valores a la sentencia return luego
	x = v * 4 / 9 //OJO que ya estan definidas, no las podemos definir nuevamente en la func
	y = v - x
	return
}

func MSum(values ...float64) float64 { //numero indeterminado de parametros, se maneja internamente como un array para luego poder accederlos
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("there are no values")
	}

	res := values[0]

	for _, v := range values[1:] {
		switch op {
		case SUM:
			res += v
		case SUB:
			res -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("zero division error")
			}
			res /= v
		case MUL:
			res *= v
		}
	}
	return res, nil
}
