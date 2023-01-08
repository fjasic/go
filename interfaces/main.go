package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb, sb := englishBot{}, spanishBot{}
	print(eb)
	print(sb)

}

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

func print(b bot) {
	fmt.Println(b.getGreeting())
}
