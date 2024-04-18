package main

import (
	"log"
	"strings"
)

func main() {
	var a strings.Builder
	// a.WriteString("\t")
	a.Write([]byte("\t"))
	res := a.String()
	log.Println(res, "111")
}
