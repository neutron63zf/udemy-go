package a

import "fmt"

// Bain a
func Bain() {
	A()
	fmt.Println("nnn")
}

// P s
type P struct {
	Name string
	age  uint8
}

// Profile s
func (p P) Profile(header string) (int, error) {
	return fmt.Println(header + p.Name)
}

func (p *P) Cname(after string) {
	p.Name = after
}
