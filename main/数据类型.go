package main

import "fmt"

func okcfunc(a, b int, c string) int {

	return a + b
}

func okcfunc1(a ...int) []int {
	fmt.Println(a)
	return a
}

func okcfunc2(a int, b ...int) []int {
	fmt.Println(a)
	fmt.Println(b)
	return b
}
func exec1() {
	a := 99
	fmt.Printf("%v -- %T\n", a, a)

}

func exec5() {
	user := make(map[string]string)
	var user1 = map[string]string{
		"name": "1",
		"b":    "2",
		"c":    "3",
	}
	user2 := map[string]string{
		"name": "11",
		"b":    "22",
		"c":    "33",
	}
	user["name"] = "yyx"
	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user["name"])
	fmt.Println(user1["name"])
	fmt.Println(user2["name"])
	for k, v := range user1 {
		fmt.Println(k, v)
	}

	fmt.Println("=== map CRUD ===")

	fmt.Println("=== C ===")
	user1["d"] = "yang"
	fmt.Println(user1)

	fmt.Println("=== C ===")
	v, ok := user1["name"]
	fmt.Println(v, ok)

	v1, ok2 := user1["xxxx"]
	fmt.Println(v1, ok2)

	fmt.Println("=== U ===")
	user1["b"] = "okc"
	fmt.Println(user1)

	fmt.Println("=== D ===")
	delete(user1, "b")
	fmt.Println(user1)

}

func exec6() {

	var userInfo = make([]map[string]string, 3, 3)
	fmt.Println(userInfo)

	var userInfo1 = []map[string]string{}
	fmt.Println(userInfo1)

	var userInfo2 = make(map[string][]string)
	userInfo2["a"] = []string{"1", "2", "3"}
	fmt.Println(userInfo2)

	aaaa := okcfunc(1, 2, "okc")
	fmt.Println(aaaa)

	funcResult1 := okcfunc1(11, 22, 33, 44, 55)
	fmt.Println(funcResult1)

	funcResult2 := okcfunc2(1, 2, 3, 4, 5, 6)
	fmt.Println(funcResult2)

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
	exec5()
}
