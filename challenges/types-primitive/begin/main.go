// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	var val int = 5

	// print the value of the local variable "val"
	fmt.Printf("local val: %T %v\n", val, val)

	// print the value of the package-level variable "val"
	fmt.Printf("global val: %T %v\n", getGlobalVal(), getGlobalVal())

	// update the package-level variable "val" and print it again
	setGlobalVal("global changed")
	fmt.Printf("global val: %T %v\n", getGlobalVal(), getGlobalVal())

	// print the pointer address of the local variable "val"
	fmt.Printf("address: %T %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 6
	fmt.Printf("val: %T %v\n", val, val)
}

func getGlobalVal() string {
	return val
}

func setGlobalVal(value string) {
	val = value
}
