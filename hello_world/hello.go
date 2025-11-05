package main

const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(who, language string) string {
	if who == "" {
		who = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + who
	} else {
		return englishHelloPrefix + who
	}
}

func main() {
}
