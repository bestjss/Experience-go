package main

import (
	"fmt"
	"study/hello"
	)
func main()  {
	fmt.Println(hello.SayHello("jS"))
	// fmt.Println(hello.SayHello(""))

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := hello.SayHellos(names)
	fmt.Println(messages,err)
}