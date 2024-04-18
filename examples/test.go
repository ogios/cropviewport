package main

import "fmt"

type A struct {
	S string
}

var a = &A{
	S: "ss",
}

func main() {
	b := []*A{nil}
	b[0] = a
	c := b[0]
	fmt.Println(c == a)
}
