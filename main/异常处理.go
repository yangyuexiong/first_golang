package main

import (
	"fmt"
)

func test0() {
	panic("抛出异常与Python的raise相同")
}

func test1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("okc")
		}
	}()
	a := 1
	b := 0
	fmt.Println(a / b)
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("没有异常")
		}
	}()
	test0()
	test1()
}
