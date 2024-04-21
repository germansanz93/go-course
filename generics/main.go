package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	v1 := []float64{1, 3, 5, 45, 12, 223, 6.92, 78.102}
	v2 := []int32{9, 23, 1, 23}

	fmt.Println(Sum01(v1))
	fmt.Println(Sum01(v2))
	fmt.Println(Sum02(v1))
	fmt.Println(Sum02(v2))
	anyType(1, 1)
	anyType02(1, "asd")
	comparableType(1, 3.23)
	orderedValues(1, 3.23)
	csInt := CustomSlice[int]{1, 2, 3, 4}
	csStg := CustomSlice[string]{"a", "b", "c"}
	fmt.Println(csInt)
	fmt.Println(csStg)
	x, y := Point(5), Point(2)
	fmt.Println(MinNumber(x, y))
	fd := MyGenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}} //Indico entre corchetes que el generico va a ser de tipo MyFirstData para esta impl
	fd.Data.PrintOne()
	sd := MyGenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}}
	sd.Data.PrintTwo()
}

func Sum01[N int | int32 | int64 | float64 | float32](v []N) N { //El generico lo defino entre corchetes, le doy un nombre (N en este caso) y le indico los valores que puede recibir
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

type Number interface {
	int | int32 | int64 | float64 | float32 //tambien podria mezclar los conceptos anteriores y definir en una interfaz que tipos acepto que se ingresen
}

func Sum02[N Number](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

func anyType[N any](v1, v2 N) { //Tambien tengo la posibilidad de usar any, sin embargo no puedo usar en este caso un int y un string por ejemplo, porque el generico tomara un solo tipo
	fmt.Println(v1, v2)
}

func anyType02[N1 any, N2 any](v1 N1, v2 N2) { //Para poder usar tipos diferentes, necesito genericos diferentes
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) { //Ojo! usamos comparable en lugar de any si queremos hacer comparaciones, esto es porque al usar any permitimos cualquier tipo de dato y hay tipos no comparables, comparable es un subconjunto de any
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
}

func orderedValues[N constraints.Ordered](v1, v2 N) { //Comparable no me permite comprar mayor o menor, solo igual o distinto, si necesito hacer esto tengo que usar una libreria exp, a partir 1.21 puedo usar cmp
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
	fmt.Println(v1 <= v2)
	fmt.Println(v1 >= v2)
}

type Point int

type CustomSlice[v int | string] []v //Creo un slice generico, contendra int o string

func MinNumber[T N1](x, y T) T { //Uso N1 y no Number para que admita Point
	if x < y {
		return x
	}
	return y
}

type N1 interface {
	~int | ~int32 | ~int64 | ~float64 | ~float32 //Con el caracter ~ antes del tipo indico que no solo acepte esos tipos, sino definiciones que terminan siendo de esos tipos tambien, como por ejemplo Point en este mismo archivo
}

type (
	MyGenericStruct[D any] struct {
		StrValue string
		Data     D
	}
	MyFirstData  struct{}
	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}
