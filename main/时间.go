package main

import (
	"fmt"
	"time"
)

func func0() {
	timeOBJ := time.Now()
	fmt.Println(timeOBJ)
	nowTimeStamp := timeOBJ.Unix()
	nowDateTime1 := timeOBJ.Format("2006-01-02 03:04:05")
	nowDateTime2 := timeOBJ.Format("2006-01-02 15:04:05")
	fmt.Println(nowTimeStamp)
	fmt.Println(nowDateTime1)
	fmt.Println(nowDateTime2)
}

func func1() {
	timeStamp := 1633584391
	toDateTime := time.Unix(int64(timeStamp), 0)
	fmt.Println(toDateTime.Format("2006-01-02 15:04:05"))

}
func func2() {
	dateTime := "2021-10-07 01:32:11"
	toTimeStamp, _ := time.ParseInLocation("2006-01-02 15:04:05", dateTime, time.Local)
	fmt.Println(toTimeStamp.Format("2006-01-02 15:04:05"))
}
func main() {
	fmt.Println(time.Now().Unix())
	func0()
	func1()
	time.Sleep(time.Second * 3)
	func2()
	fmt.Println(time.Now().Unix())
}
