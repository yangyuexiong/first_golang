package main

import "fmt"

/*
引用数据类型需要先分配内存,make或new
*/

func func3() {
	a := 33
	b := &a
	fmt.Println(a)
	*b = 99
	fmt.Println(a)
	fmt.Println(*b)
}
func main() {
	a := 33
	b := &a
	fmt.Printf("%v -- %T -- %p\n", a, a, &a)
	fmt.Printf("%v -- %T -- %p -- %v\n ", b, b, &b, *b)

	func3()

	//var user1 map[string]string  //错误
	var user1 = make(map[string]string) ////分配内存，make创建map
	user1["name"] = "okc"

	var user2 = map[string]string{}
	user2["name"] = "yyx"

	fmt.Println(user1)
	fmt.Println(user2)

	x1 := new(int) //分配内存，new创建指针
	fmt.Printf("%v -- %T -- %v", x1, x1, *x1)
}
