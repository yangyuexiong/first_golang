package main

import "fmt"

type okc1 int
type okc2 func(int, string) int

func (o okc1) show() {
	fmt.Println("相当于Python的setter")
}

type Person struct {
	name string
	age  int
}

func (p Person) show() {
	fmt.Printf("姓名:%v 年龄:%v\n", p.name, p.age)
}

// 需要使用 *Person 指针类型,否则无法修改
func (p *Person) setAge(age int) {
	p.age = age
}

type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Printf("%v  跑...\n", a.name)
}

// 结构体嵌套(继承)

type Dog struct {
	age int
	Animal
}

func (d Dog) eat() {
	fmt.Printf("%v  吃...\n", d.name)
}

func demo() {
	//实例化-方法1
	var p1 Person
	p1.name = "yangyuexiong1"
	p1.age = 18
	//fmt.Printf("%v -- %T\n", p1, p1)
	fmt.Printf("%#v -- %T\n", p1, p1)

	//实例化-方法2  结构体指针可以不使用 *号 操作对应的值
	var p2 = new(Person)
	p2.name = "yangyuexiong2"
	p2.age = 19
	//fmt.Printf("%v -- %T\n", p2, p2)
	fmt.Printf("%#v -- %T\n", p2, p2)

	//实例化-方法3 与方法2相似
	var p3 = &Person{}
	p3.name = "yangyuexiong3"
	p3.age = 20
	//fmt.Printf("%v -- %T\n", p3, p3)
	fmt.Printf("%#v -- %T\n", p3, p3)

	//实例化-方法4
	var p4 = Person{
		name: "yangyuexiong4",
		age:  21,
	}
	fmt.Printf("%#v -- %T\n", p4, p4)

	//实例化-方法5 与方法4相似,使用 & 获取地址操作值
	var p5 = &Person{
		name: "yangyuexiong5",
		age:  22,
	}
	fmt.Printf("%#v -- %T\n", p5, p5)

	//实例化-方法6 与方法4,5相似,不传key，value会使用默认值
	var p6 = &Person{
		name: "yangyuexiong6",
	}
	fmt.Printf("%#v -- %T\n", p6, p6)

	var p7 = &Person{
		"yangyuexiong6",
		99,
	}
	fmt.Printf("%#v -- %T\n", p7, p7)
}

func main() {

	var a okc1 = 1
	a.show()

	demo()

	var p = Person{
		name: "yangyuexiong4",
		age:  21,
	}
	fmt.Printf("%#v -- %T\n", p, p)
	p.setAge(9999)
	fmt.Printf("%#v -- %T\n", p, p)

	var d = Dog{
		age: 88,
		Animal: Animal{
			name: "new dog",
		},
	}
	fmt.Printf("%#v -- %T\n", d, d)
	d.run()
	d.eat()
}
