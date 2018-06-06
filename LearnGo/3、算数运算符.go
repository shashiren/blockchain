package main

import "fmt"

func main()  {
	var a  = 10
	a++
	//++a 无效
	fmt.Println(a)

	var b  = 11
	b--
	fmt.Println(b)
}

