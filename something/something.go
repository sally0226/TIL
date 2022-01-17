package something

import "fmt"

func sayBye() { // 소문자로 시작 -> private
	fmt.Println("Bye")
}

func SayHello() { // 대문자로 시작 -> public
	fmt.Println("Hello")
}
