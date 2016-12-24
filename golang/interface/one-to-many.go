// type circle has area() method to satisfy the shape interface.
// so a variable of type shape can hold a object of circle type.
// this is still true, if the circle struct has other methods that are not defined in the shape interface. in our example, circle struct has a perimeter method aswell, which is fine.
// so, for a type to implement the interface, it should atleast have the methods that the interface defines.
// this flexibility allows one type to satisfy many interfaces
// below code demonstrates that this is acceptable.

package main

import (
	"fmt"
)

type shape interface {
	area() float32
}

type shape2 interface {
	perimeter() float32
}

type circle struct {
	r float32
}

func (c circle) area() float32 {
	return 3.141 * c.r * c.r
}

func (c circle) perimeter() float32 {
	return 2 * 3.141 * c.r
}

func main() {

	var c1 shape
	c1 = circle{7}
	fmt.Println(c1)

	var c2 shape2
	c2 = circle{7}
	fmt.Println(c2)
}
