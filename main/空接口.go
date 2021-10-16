package main

import "fmt"

func show(a interface{}) {
	//println(a)
	fmt.Printf("%v %T\n", a, a)
}

func main() {
	show(99)
	show("yyx")
	show(true)

	var m1 = make(map[string]interface{})
	m1["username"] = "yyx"
	m1["age"] = 18
	m1["yyx"] = true
	m1["hobby"] = []string{"python", "golang"}
	fmt.Println(m1)

	// 错误方法
	//fmt.Println(m1["hobby"][1])

	// 正确方法 (结构同理)
	hobby2, _ := m1["hobby"].([]string)
	fmt.Println(hobby2[0])
	fmt.Println(hobby2[1])

	var s1 = []interface{}{"a", 1, true}
	fmt.Println(s1)
}
