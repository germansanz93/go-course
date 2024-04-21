package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p\n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("address: %p\n", m)
	m.Name = name
}

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println("Direccion de memoria de la var i:", &i)
	fmt.Println("Valor almacenado en iP (direccion de la var i):", iP)
	fmt.Println("Valor almacenado en i:", i)
	fmt.Println("Valor al que apunta el puntero iP finalmente (valor almacenado en var i)", *iP)
	fmt.Println("Direccion de memoria de la var iP (Recordemos que si bien es un puntero, tambien tiene su direccion en memoria para almacenar la direccion de memoria del puntero):", &iP)

	*iP = 1
	fmt.Printf("value i: %d, value pointer i: %d\n", i, *iP)

	myVar := 30
	fmt.Printf("Address %p, value %d\n", &myVar, myVar)
	increment(myVar)
	fmt.Printf("Address %p, value %d\n", &myVar, myVar)
	incrementP(&myVar)
	fmt.Printf("Address %p, value %d\n", &myVar, myVar)

	var mySlice []int

	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("address: %p, values: %v\n", mySlice, mySlice)
	fmt.Printf("addr1: %p, addr2: %p, addr3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)

	myStruct := &MyStruct{ID: 1234, Name: "Test"} //Obtengo el puntero del struct creado y lo almaceno en la var myStruct
	fmt.Printf("address: %p\n", myStruct)

	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	myStruct.Set("test method")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

	myStruct.SetP("test method 2")
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)

}

func updateSlice(v []int) { //No tiene sentido usar punteros de slicees porque siempre se usan por puntero
	fmt.Printf("address: %p\n", v)
	fmt.Printf("addr1: %p, addr2: %p, addr3: %p\n", &v[0], &v[1], &v[2]) //Notese que no cambia la direccion de memoria del slice al entrar en la funcion
	v[0] = 12                                                            //Modificar un slice en una funcion, modifica el slice original

}

func increment(val int) {
	val++
	fmt.Println(&val)
	fmt.Println(val)
}

func incrementP(val *int) {
	*val++
	fmt.Println(&val)
	fmt.Println(val)
	fmt.Println(*val)
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address in function: %p\n", v)
	v.ID = 999
	v.Name = "Xennt"
}
