package main

import "fmt"

type person struct{
	age int
	name string
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p1  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	var p = person{18, "Tom"}
	var pPoint = &p
	pPoint.age = 19
	fmt.Println("修改后的年龄为：", p.age)
	fmt.Println(v1, v2, v3, p1)
}