package main

import "fmt"

func end1() {
	fmt.Println("end1")
}

func end2() {
	fmt.Println("end2")
}

func main() {
	defer end2()
	defer end1()
	fmt.Printf("en\n")
}
