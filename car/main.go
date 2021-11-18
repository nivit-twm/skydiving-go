package main

import "fmt"

type Car struct {
	brand  string
	engine string
}

func main() {
	a := Car{"Honda", "Benzene"}
	a.printCarInfo()
	a.changeEngine("Diezel")
	// Function in golang is pass by value so when we change the value inside change engine function
	// it only change the value of the data inside that function since the function copy the data when
	// it got pass by.
	fmt.Printf("a Car in main function: %+v \n", a)
	a.changeEngineForReal("Diezel")
	fmt.Printf("a Car in main function after changeEngineForReal: %+v \n", a)
}

// Methods
// Since go don't have class so you need to define method by using a *receiver* argument
func (c Car) printCarInfo() {
	fmt.Printf("%+v \n", c)
}

func (c Car) changeEngine(newEngine string) {
	c.engine = newEngine
	fmt.Printf("car in changeEngine function %+v \n", c)
}

func (c *Car) changeEngineForReal(newEngine string) {
	c.engine = newEngine
}

// Pointer Operation in golang
// &variable -> get underlying memory address of the variable
// *pointer -> get value of the from the memory address (from the pointer)
