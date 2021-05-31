package pointers

import "fmt"

func Exercise1() {
	x := 10.10
	fmt.Printf("Address: %p", &x)
	ptr := &x
	fmt.Printf("Type: %T Value: %v", ptr, *ptr)
}

func Exercise2() {
	x, y := 10, 2
	ptrx, ptry := &x, &y
	z := *ptrx / *ptry
	fmt.Printf("%d / %d = %d", x, y, z)
}

func swap(x, y *float64) {
	*x, *y = *y, *x
}

func Exercise3() {
	x, y := 5.5, 8.8
	swap(&x, &y)
	fmt.Printf("X: %.2f Y: %.2f", x, y)
}

func main() {
	Exercise3()
}
