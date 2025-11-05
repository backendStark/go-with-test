package main

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(who, language string) string {
	if who == "" {
		who = "world"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + who
	case french:
		return frenchHelloPrefix + who
	default:
		return englishHelloPrefix + who
	}
}

func main() {
}
