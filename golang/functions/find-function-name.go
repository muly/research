package main

import (
	"fmt"
	"runtime"
)

func findFunctionName() {
	pcs := make([]uintptr, 10)
	runtime.Callers(0, pcs)
	fmt.Println()
	for i, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			fmt.Println(i, fn.Name())
		}
	}
}

func myFunc() {
	findFunctionName()
}

func init() {
	findFunctionName()
}

func main() {
	findFunctionName()
	myFunc()
}

/*
// for findFunctionName() called in init()
	0 runtime.Callers
	1 main.findFunctionName
	2 main.init.1
	3 main.init
	4 runtime.main
	5 runtime.goexit
// for findFunctionName() called in main()
	0 runtime.Callers
	1 main.findFunctionName
	2 main.main
	3 runtime.main
	4 runtime.goexit
// for findFunctionName() called in myFunc() which is called in main()
	0 runtime.Callers
	1 main.findFunctionName
	2 main.myFunc
	3 main.main
	4 runtime.main
	5 runtime.goexit
*/
