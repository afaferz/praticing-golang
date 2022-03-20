package main

import "fmt"

func main() {
	bytes, err := fmt.Println("Hello World!")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(bytes)

}
