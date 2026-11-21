package main

import "fmt"

// Greeting 返回问候语。
func Greeting(name string) string {
	if name == "" {
		name = "world"
	}
	return "hello, " + name
}

func main() {
	fmt.Println(Greeting(""))
}
