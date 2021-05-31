package strings

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Exercise1() {
	var name string = "Sergio"
	country := "Costa Rica"

	fmt.Printf("%q \n %q\n", country, name)

	fmt.Println(`He says "Hello"`)

	fmt.Println(`C:\Users\a.txt`)

}

func Exercise2() {
	r, size := utf8.DecodeRuneInString("ă")
	_ = size
	fmt.Printf("%d\n", r)

	part1, part2 := "ma", "m"
	word := part1 + part2 + string(r)
	fmt.Println(word)
}

func Exercise3() {
	s1 := "țară means country in Romanian"

	fmt.Print("Byte list: ")

	for i := 0; i < len(s1); i++ {
		fmt.Print(s1[i])
	}

	fmt.Print("\nRune list: ")

	for i := 0; i < len(s1); i++ {
		fmt.Print(utf8.DecodeRuneInString(s1[i:]))
	}
}

func Exercise4() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	s1 = strings.Replace(s1, "G", "x", 1)

	fmt.Printf("%q\n", s1)

	// printing the number of runes in the string
	fmt.Println(utf8.RuneCountInString(s1))
}

func Exercise5() {
	s := "你好 Go!"
	b := []byte(s)
	for idx, v := range b {
		fmt.Printf("Index %d Byte: %v\n", idx, v)
	}
}

func Exercise6() {
	s := "你好 Go!"
	r := []rune(s)
	for idx, v := range r {
		fmt.Printf("Index %d Rune: %v\n", idx, v)
	}
}
