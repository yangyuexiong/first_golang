package def

import "fmt"

var okc = "私有变量okc"
var Okc = "公有变量Okc"

// 公有方法
func Def(s1 string) string {
	fmt.Println("Def", s1)
	return s1
}

// 私有方法
func def(s1 string) {
	fmt.Println("def", s1)
}
