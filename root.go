package main

import (
	"fmt"

	"./pack"
)

// MyStruct s
type MyStruct struct {
	a    string
	b, c int
}

type test struct {
	email string `jsoon:"email"`
}

// MyInt u
type MyInt int

func main() {
	var g MyInt = 3
	var t test
	t.email = "abs"
	fmt.Println(g)
	fmt.Println(t)
	fmt.Println(t.email)
	var s a.P
	s.Name = "ss"
	s.Profile("hey")
	personA := a.P{
		Name: "Atsuko",
	}
	personA.Profile("how ")
	(&personA).Cname("aaaaa")
	personA.Profile("how ")
}

func main1() {
	a.Bain()
	a := [...]int{1, 2, 3, 4, 6}
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))
	a[2] = 444
	fmt.Println(a)
	/* var b []int
	b[0] = 1
	b[1] = 111
	fmt.Println(b) */
	d := map[string]int{"hello": 0, "world": 2}
	fmt.Println(d)
	fmt.Println(d["hello"])
	fmt.Println(d["world"])
}
