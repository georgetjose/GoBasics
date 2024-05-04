// Package is used as a shared Library
package main //main package is used in making the go program executable

import (
	"fmt" //fmt(format) is a package that has Print functions
)

func main() { // main function is a special function which acts as the entry point for go language
	fmt.Println("Hello World George from Within Main Function!!!")
	hello1()
	hello2()
	hello3()
}

func hello1() {
	fmt.Println("Hello World George from Outside Main Function but with same file!!!")
}

func hello3() {
	var name string
	var age int
	fmt.Println("Enter a String and a number:")
	count, err := fmt.Scanf("%s %d", &name, &age)

	fmt.Println("Count: ", count)
	fmt.Println("Error: ", err)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)

}
