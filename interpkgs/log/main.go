package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("test")

	// log.Panicln("my panic") // Ehecuta un log y  un panic, la diferencia con fatal es que panic pinta el stack trace y fatal hace el exit de inmediato

	// log.Fatal("my error") //Ejecuta un log y un exit con codigo 1

	date := time.Now()
	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
		log.Panic(err.Error())
	}

	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile) //Con log.New puedo crear un nuevo logger, al que le indico donde escribir los logs, un prefijo y un flag que indica un formato

	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")

	l.SetPrefix("errors: ") // tengo setters para modificar los valores
	l.Println("test 4")
	l.Println("test 5")

	l = log.New(os.Stdout, "success: ", log.LstdFlags|log.Lshortfile) //Si elijo stdout es la consola
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")

}
