package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// defers execution of the function until the running function end/return
	// which in this scope is main.
	defer func() {
		fmt.Println("---- run after main is done ----")
		fmt.Println("sum =", add(1, 2))
	}()
	// Declare variable
	var x int
	y := 8

	// If you declare without given the value Golang will resolve to Zero values
	// https://golang.org/ref/spec#The_zero_value
	fmt.Printf("Type of x: %T, value of x %d \n", x, x)
	fmt.Printf("Type of y: %T, value of y %d \n", y, y)

	var (
		ToBe   bool   = false
		MaxInt uint64 = 1<<64 - 1
	)
	fmt.Printf("a: %t, b: %d\n", ToBe, MaxInt)

	// const cannot be declare using :=
	const Pi = 3.14

	// Type conversion
	fmt.Printf("type after casting int to float64 %T \n", float64(y))

	// Control Statements
	// For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// Golang dont' have while loop! we used for w/o any condition to replace that
	// for sum < 1000 {
	//
	// }

	// If (short statement)
	if _, err := os.ReadFile("not_exist.txt"); err != nil {
		fmt.Println(err)
	}

	// We don't need break statement in golang,
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch w/out condition can be found a lot when you have to deal with channel
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("----------------------------")

	// Struct
	type car struct {
		brand  string
		engine string
		fuel   int
	}

	var carA car
	carA.brand = "Honda"
	carA.engine = "Diesel"
	carA.fuel = 70

	carB := car{"Honda", "Diesel", 70}
	fmt.Printf("%+v \n", carB)
	carB.engine = "Benzene"
	fmt.Printf("%+v \n", carB)
}

// function w/ named return
func add(x, y int) (sum int) {
	sum = x + y
	return
}
