package main

import "fmt"

func main()  {
	const(
		a=iota
		b=iota
		c=iota
		d=iota
	)
	fmt.Println(a,b,c,d)//统计一行代码中被调用的次数
}

