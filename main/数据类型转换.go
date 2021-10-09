package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func getByte() {
	// 获取字节大小，内存大小，1M = 1024k = 1048576字节
	fmt.Println("=== int8,int64 ===")
	var a1 = 1
	var a2 int
	a2 = 2
	var a3 int8 = 1
	var a4 int64 = 1
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(unsafe.Sizeof(a1))
	fmt.Println(unsafe.Sizeof(a2))
	fmt.Println(unsafe.Sizeof(a3))
	fmt.Println(unsafe.Sizeof(a4))

	fmt.Println("=== string ===")
	var b1 = "okc"
	var b2 string
	b2 = "okc2"
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(unsafe.Sizeof(b1))
	fmt.Println(unsafe.Sizeof(b2))

	fmt.Println("=== float32, float64 ===")
	var f1 float32 = 1.23
	var f2 float64 = 1.23
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(unsafe.Sizeof(f1))
	fmt.Println(unsafe.Sizeof(f2))
}

func toStrTest() {
	// 转字符串
	i := 33
	f := 33.33
	b := true
	by := 'a'

	to1 := fmt.Sprintf("%d", i)
	to2 := fmt.Sprintf("%f", f)
	to3 := fmt.Sprintf("%t", b)
	to4 := fmt.Sprintf("%c", by)
	fmt.Printf("%v %T\n", to1, to1)
	fmt.Printf("%v %T\n", to2, to2)
	fmt.Printf("%v %T\n", to3, to3)
	fmt.Printf("%v %T\n", to4, to4)

	// 转整数
	// 转浮点数
	// 转bool
	// 转
	res1 := strconv.FormatInt(int64(i), i)
	res2 := strconv.FormatFloat(float64(f), 'f', 2, 64)
	res3 := strconv.FormatBool(b)
	res4 := strconv.FormatUint(uint64(by), 10)
	fmt.Printf("%v %T\n", res1, res1)
	fmt.Printf("%v %T\n", res2, res2)
	fmt.Printf("%v %T\n", res3, res3)
	fmt.Printf("%v %T\n", res4, res4)

	sss := "123456789"
	sss1 := "123.456"
	res5, _ := strconv.ParseInt(sss, 10, 64)
	res6, _ := strconv.ParseFloat(sss1, 64)
	fmt.Printf("%v %T\n", res5, res5)
	fmt.Printf("%v %T\n", res6, res6)
}
func main() {
	getByte()
	toStrTest()
}
