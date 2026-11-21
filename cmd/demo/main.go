package main

import (
	"fmt"
	"strings"
)

// Greeting 返回问候语。
func Greeting(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "world"
	}
	return "hello, " + name
}

func main() {
	fmt.Println(Greeting(""))
}
