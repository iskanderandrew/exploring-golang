package main

import "fmt"

var a2 int
type boy int
var b2 boy

func main()  {
	a2 = 42
	b2 = 43
	fmt.Println(a2)
	fmt.Printf("%T\n", a2)
	fmt.Println(b2)
	fmt.Printf("%T\n", b2)
	a2 = int(b2)
	fmt.Println(a2)
	fmt.Printf("%T\n", a2)
}
