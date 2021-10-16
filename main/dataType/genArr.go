package main

import (
	"fmt"
	"strings"
)

func genArr() {
	// 创建数组
	arr := []string{"java", "python", "golang"}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	arr2 := strings.Join(arr, "---")
	fmt.Println(arr2)

	// 集合 str1 包含 str2
	str1 := "yang yuei xiong"
	str2 := "yang"
	flag := strings.Contains(str1, str2)
	println(flag)
}

func main() {
	genArr()
}
