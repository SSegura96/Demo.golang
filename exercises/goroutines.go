package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(x float32, y float32, wg *sync.WaitGroup) {
	fmt.Printf("Total: %.2f\n", x+y)
	wg.Done()
}

func deposit(b *int, n int, wg *sync.WaitGroup, mux *sync.Mutex) {
	mux.Lock()
	*b += n
	mux.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mux *sync.Mutex) {
	mux.Lock()
	*b -= n
	mux.Unlock()
	wg.Done()
}

func Execise1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("Mr. Wick", &wg)
	wg.Wait()
}

//Perform 3 additions async
func Execise2() {
	var wg sync.WaitGroup
	wg.Add(3)
	go sum(1, 2, &wg)
	go sum(3, 1, &wg)
	go sum(20, 5, &wg)
	wg.Wait()
}

// Calculate the square root with a anon function
func Execise3() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(f int, wg *sync.WaitGroup) {
		r := math.Sqrt(float64(f))
		fmt.Printf("The square root of %d is %.2f", f, r)
		wg.Done()
	}(9, &wg)
	wg.Wait()
}

func Execise4() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100.; i <= 149; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			r := math.Sqrt(f)
			fmt.Printf("The square root of %.0f is %.2f\n", f, r)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

func Execise5() {
	var wg sync.WaitGroup
	var mux = sync.Mutex{}

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mux)
		go withdraw(&balance, i, &wg, &mux)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

func main() {
	Execise5()
}
