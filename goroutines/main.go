package main

import (
	"fmt"
	"time"
)

func main() {
	go myProcess("A")
	go myProcess("B")
	time.Sleep(2 * time.Second)

	myFirstChannel := make(chan string) //Creo un canal

	go myProcessWithChannel("C", myFirstChannel) //Recibo el mensaje del canal

	result := <-myFirstChannel //Tomo el valor de mi canal, si no hago esto el programa finalizara sin esperar que la goroutine donde esta corriendo el proceso finalice
	fmt.Println(result)

	close(myFirstChannel) //Cuando abro un canal, luego al final de mi ejecucion debo cerrrarlo con la palabra reservada close

	channelA := make(chan string)
	channelB := make(chan string)

	go myProcessWithChannel("D", channelA)
	go myOtherProcessWithChannel("E", channelB)

	resultA := <-channelA //OJO ACA, este me esta bloqueando continunar con el codigo hasta que llegue un mensaje a este canal
	fmt.Println("A: ", resultA)
	resultB := <-channelB //OJO ACA, porque si bien este ejecuta solo con 5 valores y termina primero que el channelA que tiene 20, estamos a la espera de resultA para continuar y por lo tanto esto esta en pausa
	fmt.Println("B: ", resultB)
	close(channelA)
	close(channelB)
}

func myOtherProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 5 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process %s - num %d\n", p, i)
	}
	c <- "ok" //envio un mensaje al canal
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process %s - num %d\n", p, i)
	}
	c <- "ok" //envio un mensaje al canal
}

func myProcess(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process %s - num %d\n", p, i)
	}
}
