package main

import "fmt"

// Pai is test const
const Pai float32 = 3.14
const (
	// ZERO is test const
	ZERO = iota
	ONE
	TWO
	THREE
	FOUR
)

func main() {
	var p *int
	n := 10
	p = &n
	fmt.Println(p)
	fmt.Println(*p)
	// *int つけるとlintが怒る
	var h = new(int)
	fmt.Println(*h)
	fmt.Printf("%v\n", Pai)
	fmt.Printf("%v %v %v %v %v\n", ZERO, ONE, TWO, THREE, FOUR)
}
