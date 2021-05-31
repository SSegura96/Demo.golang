package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
	model     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return fmt.Sprint(c.brand, " ", c.model)
}

func Exercise1() {
	var mustang vehicle = car{licenceNo: "123456", brand: "Ford", model: "Mustang"}
	fmt.Println(mustang.License())
	fmt.Println(mustang.Name())
}

func Exercise2() {
	var empty interface{}
	fmt.Printf("%T", empty)
	empty = 1
	empty = 16.
	empty = []int{1, 2, 3, 4}
	empty = append(empty.([]int), 5)
	fmt.Printf("%#v", empty)
}

func main() {
	Exercise2()
}
