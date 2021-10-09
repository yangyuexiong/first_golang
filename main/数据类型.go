package main

import (
	"fmt"
)

func exec1() {
	a := 99
	fmt.Printf("%v -- %T\n", a, a)
}

func exec2() {
	str1 := "yang yuei xiong"
	arr1 := []string{"java", "python", "golang"}
	for _, i := range str1 {
		fmt.Printf("%c %v %T\n", i, i, i)
	}

	for _, i := range arr1 {
		fmt.Printf("%v %T\n", i, i)
	}
}

func exec3() {

	// 字符串替换 ascii，utf8
	y := "yyx"
	byteStr1 := []byte(y)
	byteStr2 := []rune(y)
	byteStr1[0] = 'x'
	byteStr2[0] = '杨'
	fmt.Println(string(byteStr1))
	fmt.Println(string(byteStr2))

}

var x1 = 1
var s1 = "yangyuexiong"

const x2 = 999
const s2 = "yyx"

func main() {
	println("hello world")
	print("hello golang\n")
	fmt.Printf("%v -- %T\n", x1, x1)
	fmt.Printf("%v -- %T\n", s1, s1)
	fmt.Printf("%v -- %T\n", x2, x2)
	fmt.Printf("%v -- %T\n", s2, s2)
	exec1()
	exec2()
	exec3()

}
