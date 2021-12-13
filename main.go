package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer	func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Error occured")
	fmt.Println("End")
}
