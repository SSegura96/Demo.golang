package main

import "fmt"

const iva = 0.13
const sale = 0.10

type money float64

type book struct {
	title  string
	author string
	price  float64
}

func (m money) print() {
	fmt.Printf("%.2f", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func (b book) getIVA() float64 {
	return b.price * iva
}

func (b book) toString() {
	fmt.Printf("%#v\n", b)
}

func (b book) discont() float64 {
	return b.price - (b.price * sale)
}

func Exercise1() {
	var crc money = 1000
	crc.print()
}

func Exercise2() {
	var crc money = 1000
	fmt.Println(crc.printStr())
}

func Exercise3() {
	var antidote = book{title: "the antidote", author: "Oliver Burkeman", price: 10000.}
	antidote.toString()
	fmt.Printf("IVA: %.2f\n", antidote.getIVA())
	fmt.Printf("With discont: %.2f\n", antidote.discont())
}

func main() {
	Exercise3()
}
