package main

import "fmt"

type math interface {
	isEven(x int) bool
}
type basicmath struct {
	name string
}
type algebra struct {
	name string
	Id   int
}

func (b basicmath) isEven(x int) bool {
	fmt.Println(b.name)
	if x%2 == 0 {
		return true
	} else {
		return false
	}

}
func (c algebra) isEven(y int) bool {
	fmt.Println(c.name)
	if y == 11 {
		return true
	} else {
		return false
	}

}
func main() {
	var m math
	var n math
	m = basicmath{"my basic math"}
	n = algebra{"my algebra math", 5}
	//n.id = 5
	num := 11
	num2 := 11
	fmt.Println(num, m.isEven(num))
	fmt.Println(num2, n.isEven(num2), n.Id)
}
