package main

import (
	"fmt"
	say_hello "github.com/jeagercoder/test-my-module-go"
)


func main() {
	fmt.Println("main")
	hello := say_hello.SayHelloMamang()
	fmt.Println(hello)
}