package main

import (
	"fmt"
)

func main() {
	var a int64
	a = int64(300)

	var f float32
	f = float32(0.007)

	var c complex64

	fmt.Println("簡単なテストプログラム")
	fmt.Printf("%d", a)
	fmt.Println("　")
	fmt.Printf("%f", f)
	fmt.Println("　")
	fmt.Printf("%f", c)

}
