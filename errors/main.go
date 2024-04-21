package main

import (
	"errors"
	"fmt"

	"github.com/germansanz93/go-course/errors/mypackage"
)

type MyCustomErr struct {
	ID string
}

func (m MyCustomErr) Error() string { //Metodo que implementa la interfaz error
	return fmt.Sprintf("error with id: %s", m.ID)
}

func main() {
	var err error
	err = errors.New("my new error")
	fmt.Println(err)
	fmt.Println(err.Error()) //El metodo Error de un error nos devuelve un string con el mensaje del error

	err2 := fmt.Errorf("my format err, string: %s, number: %d", "my string", 1) //El paquete fmt tambien nos permite generar errores, sin embargo con Errorf los podemos crear formateados como hacemos con Printf y los strings
	fmt.Println(err2.Error())

	err3 := TestError(1)
	fmt.Println(errors.As(err3, &MyCustomErr{})) //En este caso el error se genero a partir de la estructura, por lo tanto responde true

	err4 := TestError(4)
	fmt.Println(errors.As(err4, &MyCustomErr{})) //No se genero a partir de la estructura, responde false
	fmt.Println(err4)

	fmt.Printf("\n\n")

	err5 := errors.Join(err, err2, err4) //Genero un errpr a partir de varios errores
	fmt.Println(err5)

	fmt.Println(errors.Is(err5, err4)) //estoy preguntando si err5 contiene a err4
	fmt.Println(errors.Is(err5, err3)) //estoy preguntando si err5 contiene a err3

	err6 := fmt.Errorf("error6: [%w]", err) //errores anidados
	err7 := fmt.Errorf("error7: [%w]", err6)

	fmt.Println(err7)
	fmt.Println(errors.Unwrap(err7))                               //obtengo el error anidado
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err7)))) // En el caso de no haber mas errores anidados, obtengo nil

	defer func() { //Defer siempre se ejecuta luego de la ejecucion
		fmt.Println("End main")
		r := recover() //Con recover capturos los errores que se hayan podido generar
		if r != nil {
			fmt.Println("Recover in", r)
		}
	}()
	// v := 0
	// _ = 5 / v            //se genera un error
	// panic("my panic")
	mypackage.Run()
	fmt.Println("end 2") // esta linea nunca se ejecuto porque el programa finalizo antes por el error
}

func TestError(v int) error {
	if v == 1 {
		return MyCustomErr{
			ID: "4",
		}
	} else {
		return errors.New("my other error")
	}

}
