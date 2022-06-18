package main

import (
	"fmt"
	"go-playground/apps/api/handlers"
)

func Hello(name string) string {
	result := "Hello, " + name
	return result
}

func main() {
	fmt.Println(Hello("api"))
	handlers.Example()
}
