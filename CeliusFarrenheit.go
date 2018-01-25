package main

import (
	"fmt"
	"strconv"
)


//°C = (°F - 32) * 5/9
//(von Fahrenheit nach Celsius )
//°F = °C * 1,8 + 32
//(von Celsius nach Fahrenheit)

var fahr, cels int

func formel(ein int) int {
	ret := (ein - 32) * 5 / 9
	return ret
}

func main() {
	bereich :=[]int{1,2,3,4,5,7}
	for _, init := range bereich {
	x := formel(init)
	ausgabe :=strconv.Itoa(x)
	fmt.Printf("%s %s\n","In Celsius: ",ausgabe)}
}
