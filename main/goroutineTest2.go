package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func ttt(n int) {
	for num := (n - 1) * 30000; num < n*30000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
			if flag {
				//fmt.Println(num)
			}
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()

	//for num := 2; num < 100000; num++ {
	//	var flag = true
	//	for i := 2; i < num; i++ {
	//		if num%i == 0 {
	//			flag = false
	//			break
	//		}
	//		if flag {
	//			//fmt.Println(num)
	//		}
	//	}
	//}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go ttt(i)
	}
	wg.Wait()
	end := time.Now().Unix() - start
	fmt.Println(end)

}
