package main

import (
	"fmt"
	"strconv"
)
func main1()  {
	var num  = 18.5
	fmt.Println(num)
	fmt.Println(int(num))
}
func main2() {
	var mystr = "12345"
	//var num = int(mystr)
	num, error :=strconv.Atoi(mystr)
	num-=5
	fmt.Println(mystr,num,error)
}

func main()  {
	var num = 12345
	mystr:=strconv.Itoa(num)
	fmt.Println(mystr)
	fmt.Println(mystr+"abc")
}