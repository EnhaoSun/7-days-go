package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (rt Rectangle) Area() float32 {
	return rt.length * rt.width
}

func main() {
	/*
	   多态的Go版本:
	   结构体 Square 实现了接口 Shaper
	   所以可以将一个 Square 类型的变量赋值给一个接口类型的变量：areaIntf = sq
	*/
	sq := new(Square)
	sq.side = 5
	var areaIntf Shaper
	areaIntf = sq
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	// 扩展一下上面的例子
	sq2 := &Square{3}
	rt := Rectangle{5, 3}

	shapes := []Shaper{sq2, rt}

	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape: ", shapes[n].Area())
	}

}
