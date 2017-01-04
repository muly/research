package main

type myInterface interface {
	IsValid() bool
}

type myFuncType func() bool

type myFuncType1 func() bool

func (k myFuncType) IsValid() bool {
	return k()
}

func (k myFuncType1) IsValid1() bool {
	return k()
}

func main() {
	// quick and dirty method:
	//	if below compiles sucessfully then the type `myFuncType` implement the `myInterface` interface
	//  source: http://goblog.qwest.io/2016/05/functions-can-implement-interface-in.html
	var _ myInterface = (*myFuncType)(nil)  // no error
	var _ myInterface = (*myFuncType1)(nil) // ERROR: *myFuncType1 does not implement myInterface (missing IsValid method)

}
