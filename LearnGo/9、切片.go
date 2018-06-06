package main

import "fmt"

func main()  {

	var numbers []int
	printSlice(numbers)

	//允许追加空切片
	numbers = append(numbers,0)
	printSlice(numbers)

	//向空切片添加一个元素
	numbers = append(numbers,1)
	printSlice(numbers)

    //同时添加多个元素
	numbers = append(numbers,2,3,4)
	printSlice(numbers)

	//创建切片 number1 是之前 切片容量的两倍,容量的值只有1,2,4,6,8
	numbers1:= make([]int,len(numbers),(cap(numbers))*2)

	//拷贝 number 的内容到number1
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice (x []int){

	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}