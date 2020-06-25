package main

import "fmt"

func main() {
	fmt.Println("Hellow everone")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Yo")
		}
		fmt.Println(i)
	}
	bar()
}

func foo() {
	fmt.Println("Im in foo")
}
func bar() {
	fmt.Println("and then we exited")
}