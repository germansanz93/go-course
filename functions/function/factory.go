package function

// Notese que no hice import de Operation y lo puedo usar igual, esto es porque estan en el mismo package
func FactoryOperation(op Operation) func(float64, float64) float64 { //Notese que se esta indicando como retorno una func que toma dos float64 y retorna un float64
	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case DIV:
		return div
	case MUL:
		return mul
	}
	return nil
}

func sub(a, b float64) float64 {
	return a - b
}

func sum(a, b float64) float64 {
	return a + b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
