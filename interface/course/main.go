package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

var _ bot = (*englishBot)(nil)
var _ bot = (*spanishBot)(nil)

func main() {
	en := englishBot{}
	sp := spanishBot{}
	printGreeting(en)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "¡Hola!"
}
