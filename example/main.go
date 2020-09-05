package main

import "log"

func main() {
	log.Println("hello oce")
	a := map[string]string{
		"a": "A",
	}
	b, c := a["a"]
	log.Printf("b: %v, c: %v", b, c)
}
