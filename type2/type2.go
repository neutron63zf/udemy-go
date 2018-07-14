package main

import "fmt"

// Sample this is sample type
type Sample struct {
	X int
	Y int
}

func main() {
	var s Sample
	fmt.Print(s.X)
	fmt.Print(s.Y)
}
