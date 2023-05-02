package main

import (
	"fmt"
	say_hello "github.com/jeagercoder/test-my-module-go"
)


func main() {
	fmt.Println("main")
	str := say_hello.SayHelloMamang()
	fmt.Println(str)
}