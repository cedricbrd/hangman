package main

import "fmt"

type turtle struct {
	name string
	age  int
}

func main() {
	var t turtle
	t.name = "oui"
	SetTurtleName("non", &t)
	PrintTurtleName(&t)
}

func PrintTurtleName(t *turtle) {
	fmt.Printf(t.name)
}

func SetTurtleName(name string, t *turtle) {
	t.name = name
}
