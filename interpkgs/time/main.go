package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() //Me devuelve la fecha actual

	p(now)

	then := time.Date(2020, 11, 17, 20, 34, 58, 651387237, time.UTC) //Creo una fecha determinada

	p(then)

	then = then.Add(1 * time.Hour) //puedoo sumar tiempo

	p(then)

	then = then.Add(24 * time.Hour) //puedoo sumar tiempo

	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p("Then is before now:", then.Before(now))
	p("Then is after now:", then.After(now))
	p("Then is now:", then.Equal(now))

	diff := now.Sub(then) //puedo restar tiempos

	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
}
