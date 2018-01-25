package main

import "fmt"

var Zeichenketten struct{
	Satz1 string
	Satz2 string
}
var Nummern struct{
	Zahl1 int64
	Zahl2 int64
}
type I interface{}

func main() {

	input := Zeichenketten
	input.Satz2 ="BOBO"
	fmt.Scan(&input.Satz1)
	fmt.Print(input)

}