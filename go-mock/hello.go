package main

import (
	"fmt"
)

// Constants can be good in terms of capturing values and aiding performance.
const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}