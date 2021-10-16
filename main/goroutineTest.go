package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
并发：
1.多个线程同事竞争一个位置，竞争到了才可以执行，每一个时间段只有一个线程在执行，
2.多个线程作用在一个cpu上面，同一时间【点】，只有一个线程执行，同一时间【段】可以又多个线程在执行
*/

// 并行：多个线程可以同时执行，每一个时间段，可以有多个线程同时执行

//var wg sync.WaitGroup

func t() {
	for i := 0; i < 10; i++ {
		fmt.Println("okc -> t()", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func t1(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, i)

	}
}
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	runtime.GOMAXPROCS(cpuNum - 2)
	fmt.Println(cpuNum)

	wg.Add(1)
	go t() // 开始协程
	for i := 0; i < 10; i++ {
		fmt.Println("okc -> main()", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("完成")
}
