package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int, c bool) int {
	fmt.Println(c)
	return a * b
}
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func lenAndUpper2(name string) (length int, uppercase string) { //naked function
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
func repeatMe(words ...string) {
	fmt.Println(words)
}
func main() {
	// fmt.Println(multiply(2, 10, true))
	// totalLength, upperName := lenAndUpper("bada")
	// fmt.Println(totalLength, upperName)
	repeatMe("bada", "daisy", "summer")
}
