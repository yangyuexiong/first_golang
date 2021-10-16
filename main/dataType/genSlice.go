package main

import (
	"fmt"
	"sort"
)

func genSlice() {

	// 创建切片
	var arr1 [3]int
	arr1[0] = 8
	arr1[1] = 7
	//var arr2 [4]int
	var arr3 [3]string
	arr3[0] = "python"

	var arr4 = [3]int{1, 2, 3}
	arr5 := [3]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5}
	arr7 := [...]int{0: 99, 1: 88, 2: 77, 3: 66, 10: 9999}
	fmt.Printf("值:%v 类型:%T\n", arr1, arr1)
	//fmt.Printf("%v %T\n", arr2, arr2)
	fmt.Printf("值:%v 类型:%T\n", arr3, arr3)
	fmt.Printf("值:%v 类型:%T\n", arr4, arr4)
	fmt.Printf("值:%v 类型:%T\n", arr5, arr5)
	fmt.Printf("值:%v 类型:%T\n", arr6, arr6)
	fmt.Printf("值:%v 类型:%T\n", arr7, arr7)

	// len长度
	// cap容量
	// make
	// 下标不能扩容，需要用append() 每次扩容翻倍,四分之一
	var sliceA = make([]int, 4, 8)
	var sliceB = make([]int, 4, 8)
	//var sliceC = make([]int, 4, 8)
	fmt.Printf("值:%v 类型:%T 长度:%v 容量:%v\n", sliceA, sliceA, len(sliceA), cap(sliceA))
	sliceA[0] = 11
	sliceA[1] = 22
	sliceA[2] = 33
	sliceA[3] = 44
	sliceB = append(sliceB, 100, 100)
	//sliceC = append(sliceC, 200, 200)
	sliceA = append(sliceA, sliceB...)
	fmt.Printf("值:%v 类型:%T 长度:%v 容量:%v\n", sliceA, sliceA, len(sliceA), cap(sliceA))
	// 升序
	sort.Ints(sliceA)
	fmt.Println(sliceA)
	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(sliceA)))
	fmt.Println(sliceA)

}

func main() {
	genSlice()
}
