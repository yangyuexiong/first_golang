package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 11
	ch <- 12
	<-ch
	a := <-ch
	fmt.Println(a)
	fmt.Printf("类型：%T 值：%v 容量：%v 长度：%v", ch, ch, cap(ch), len(ch))
}
