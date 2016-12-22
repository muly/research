//from specs:
// The method set of any other type (i.e. other than interface) T consists of all methods declared with receiver type T.
// The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T
// (that is, it also contains the method set of T).

package main

import (
	"fmt"
)

type circle struct {
	r float32
}

func (c circle) area() float32 {
	return 3.141 * c.r * c.r
}

func (c *circle) areaPointerReceiver() float32 {
	return 3.141 * c.r * c.r
}

func main() {
	c := circle{7}
	p := &circle{7}

	fmt.Println("area =", c.area())                                    // working
	fmt.Println("area =", c.areaPointerReceiver())                     // working: per method sets defined in specs, I thought this should error out, but turns out that I'm wrong. see http://stackoverflow.com/questions/33587227/golang-method-sets-pointer-vs-value-receiver, http://stackoverflow.com/questions/19433050/go-methods-sets-calling-method-for-pointer-type-t-with-receiver-t
	fmt.Println("area =", p.area())                                    // working
	fmt.Println("area =", p.areaPointerReceiver())                     // working
	fmt.Println("area =", createCirclePointer().area())                // working
	fmt.Println("area =", createCirclePointer().areaPointerReceiver()) // working
	fmt.Println("area =", createCircle().area())                       // working
	fmt.Println("area =", createCircle().areaPointerReceiver())        // ERRORs: "cannot call pointer method on createCircle()", "cannot take the address of createCircle()"

}

func createCircle() circle {
	return circle{7}
}

func createCirclePointer() *circle {
	return &circle{7}
}
