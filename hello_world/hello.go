package main

const englishHelloPrefix = "Hello, "

func Hello(who string) string {
	if who == "" {
		return englishHelloPrefix + "world"
	} else {
		return englishHelloPrefix + who
	}
}

func main() {
}
