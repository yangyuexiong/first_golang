package main

// 与python的抽象基类相似
type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话,必须要通过结构体或者自定义类型实现这个接口

type Phone struct {
	Name string
}

func (p Phone) start() {
	println(p.Name + "执行 start 方法")
}

func (p Phone) stop() {
	println(p.Name + "执行 stop 方法")
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

func main() {

	p := Phone{
		"iPhone13",
	}
	p.start()
	p.stop()

	var p1 Usber = p //golang中接口是一个数据类型
	//p1 = p       // 接口实现
	p1.start()
	p1.stop()

	var computer = Computer{}
	var phone = Phone{
		"iPhone14",
	}
	computer.work(phone)
}
