package main

import (
	"fmt"
)

/* type Hello struct {
	var aa int =0
} */

func main() {
	var a int64
	a = int64(300)

	var f float32
	f = float32(0.007)

	var c = 1 + 2i
	d, e := "hello", 22

	fmt.Println("簡単なテストプログラム")
	fmt.Printf("%d", a)
	fmt.Println("")
	fmt.Printf("%f", f)
	fmt.Println("")
	fmt.Printf("%f", c)
	fmt.Println("")
	fmt.Printf("%v %v", d, e)
	fmt.Print("\n")

}
