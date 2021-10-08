package main

import (
	"encoding/json"
	"fmt"
)

/*
首字母大写：公有
首字母小写：私有
*/

type Role struct {
	Uid []int
}
type User struct {
	ID   int `json:"id"` // 设置转后的key
	Name string
	Role
}

func main() {

	var u = User{
		ID:   33,
		Name: "yyx",
		Role: Role{
			Uid: []int{1, 2, 3},
		},
	}
	fmt.Printf("%#v -- %T\n", u, u)
	jsonByte, _ := json.Marshal(u)
	fmt.Println(jsonByte)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

	var dict = `{"ID":33,"Name":"yyx"}`
	fmt.Println(dict)
	var u1 User
	err := json.Unmarshal([]byte(dict), &u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u1)
	fmt.Printf("%v %T\n", u1, u1)
	fmt.Printf("%#v %T\n", u1, u1)
}
