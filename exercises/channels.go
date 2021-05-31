package main

import "fmt"

func power(num int, ch chan int) {
	ch <- num * num
}

func Exercise1() {
	//bidirectional unbuffered
	var c1 chan float64
	_ = c1

	//Recives - only
	c2 := make(<-chan rune)
	_ = c2

	//Send - only
	c3 := make(chan<- rune)
	_ = c3

	//With a capacity of 10 ints
	c4 := make(chan int, 10)

	fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)
}

func Exercise2() {
	ch := make(chan string)

	//a.k.a Anon function
	go func(text string) {
		ch <- text
	}("Sergio")

	value := <-ch
	fmt.Println(value)
}

func Exercise3() {
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}

func Exercise4() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}

}

func Exercise5() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go func(num int) {
			ch <- num * num
		}(i)
	}
	for i := 1; i <= 50; i++ {
		fmt.Printf("%d\n", <-ch)
	}
}

func main() {
	Exercise5()
}
