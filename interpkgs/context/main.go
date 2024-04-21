package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() //Context nos devuelve un contexto, lo podemos usar para mandarle diferentes valores y tenerlos disponibles en diferentes capas de nuestra app
	ctx = context.WithValue(ctx, "my-key", "my value")
	ctx = context.WithValue(ctx, "my-key-int", 5)
	viewContext(ctx)

	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second) //Me permite definirle un timeout a un contexto, ademas del contexto me devuelve una funcion cancel que debo utilizar para eliminarlo

	defer cancel() //Cuando se termine la ejecucion del main llamo a cancel

	go myProcess(ctx2)

	select {
	case <-ctx2.Done():
		fmt.Println("we've exceeded the deadline")
		fmt.Println(ctx2.Err())
	}
}

func viewContext(ctx context.Context) {
	fmt.Printf("my value is: %s\n", ctx.Value("my-key"))
	fmt.Printf("my other value is: %d\n", ctx.Value("my-key-int"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ohh, time out")
			return
		default:
			fmt.Println("doing some process")
		}
		time.Sleep(500 * time.Millisecond)

	}
}
