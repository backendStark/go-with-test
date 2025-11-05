package main

import "fmt"

func Hello(who string) string {
	if who == "" {
		return "Hello, world"
	} else {
		return fmt.Sprintf("Hello, %s", who)
	}
}

func main() {
}
