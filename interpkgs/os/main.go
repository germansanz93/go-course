package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.go") //Con open puedo abrir ficheros del sistema
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //Exit finaliza el programa, 0 significa que finalizo ok y 1 que finalizo con errores
	}

	fmt.Println(file)
	v, _ := file.Stat() //Stat me da metodos que puedo usar para obtener informacion del archivo
	fmt.Println(v.Name())
	fmt.Println(v.Size())

	data := make([]byte, 2048) //Creo un buffer para leer el archivo
	c, _ := file.Read(data)    //c es la cantidad de bytes del file
	fmt.Println(c)
	fmt.Println(data)                               //Al imprimir vemos si se ocupo completamente o no, ya que de no ocuparse por completo, habra ceros al final.
	fmt.Println(data[:c])                           //asi le solo lo que no es cero
	fmt.Println(string(data[:c]))                   //Obtengo el texto como tal y no un array de bytes
	fmt.Printf("read: %d bytes: %q\n", c, data[:c]) //se printea mostrando los saltos de linea en lugar de interpretarlos

	fmt.Println()
	fmt.Println()

	fmt.Println("Test variable: ", os.Getenv("TESTVAR")) //obtiene variables de entorno, en caso de bash solo las exportadas
	os.Setenv("NEW_ENV", "new_env")                      //Crea variables de entorno
	fmt.Println("Test variable: ", os.Getenv("NEW_ENV")) //obtiene variables de entorno, en caso de bash solo las exportadas

	dir, _ := os.Getwd() //obtengo el working directory
	fmt.Println(dir)

	os.Setenv("ENV_EXISTS", "") //Una variable que existe pero no tiene ningun valor
	fmt.Println(os.Getenv("ENV_EXISTS"))
	fmt.Println(os.Getenv("ENV_NOT_EXISTS"))

	env, ok := os.LookupEnv("ENV_EXISTS") //LookUpEnv nos devuelve 2 valores, la variable en custion y true o false dependiendo de si existe o no
	fmt.Println(env, ok)
	env, ok = os.LookupEnv("ENV_NOT_EXISTS")
	fmt.Println(env, ok)

	os.Setenv("DB_USERNAME", "gsanz")
	os.Setenv("DB_PASSWORD", "mypassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "mydatabase")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT") //ExpandEnv nos formatea las variables de entorno, notese que puedo indistintamente usar llave o no
	fmt.Println(dbURL)

}
