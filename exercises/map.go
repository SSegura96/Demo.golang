package map

import "fmt"

func Exercise1() {
	var m1 map[string]int
	_ = m1

	m2 := map[int]string{0: "Sergio", 1: "Sevilla"}
	m2[10] = "Abba"

	fmt.Printf("Existing key: %q Non Existing key: %q", m2[10], m2[15])

}

func Exercise2() {
	var m1 = map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	s2 := fmt.Sprintf("%d", m2)
	s3 := fmt.Sprintf("%d", m3)

	fmt.Println(s2 == s3)
}

func Exercise3() {
	m := map[int]bool{1: true, 2: false, 3: false}
	_ = m
	fmt.Printf("%#v", m)
}
