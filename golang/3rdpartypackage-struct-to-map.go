package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type Server struct {
	Name    string
	ID      int32
	Enabled bool
}

func main() {
	s := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	m := structs.Map(s)

	fmt.Println(m)

}
