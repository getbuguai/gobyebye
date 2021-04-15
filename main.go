package main

import "fmt"

func main() {
	fmt.Println("Hello Get BuGuai !!!")
	fmt.Println("----------------")
	err := Clean()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("----------------")
	fmt.Println("Clean over ...")
}
