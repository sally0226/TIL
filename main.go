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
	defer fmt.Println("I'm done") // 함수가 종료되면 실행됨
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
func repeatMe(words ...string) {
	fmt.Println(words)
}
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // idx, number
		total += number
	}
	return 1
}
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // if에서 for처럼 변수를 만들 수 있음
		return false
	}
	return true
}
func canIDrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
func main() {
	// fmt.Println(multiply(2, 10, true))
	// totalLength, upperName := lenAndUpper("bada")
	// fmt.Println(totalLength, upperName)
	// repeatMe("bada", "daisy", "summer")
	// superAdd(10, 2, 3, 4, 5, 6)
	// fmt.Println(canIDrink2(16))

	a := 2

}
