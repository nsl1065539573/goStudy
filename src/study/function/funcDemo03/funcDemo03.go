package main

import "fmt"

func main() {
	// _, x := test()
	// fmt.Println(x)
	// sum, avg := test2(2, 4)
	// fmt.Printf("sum is %v, avg is %v", sum, avg)
	x := test4(test2(2, 4))
	fmt.Printf("可以被其他函数调用%v", x)
}

// Go的函数可以声明多个返回值，接受者如果不需要可以使用"_"来忽略对应位置的返回值
func test() (int, int) {
	return 1, 2
}

// Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
func test2(x, y int) (sum int, avg int) {
	sum = x + y
	avg = sum / 2
	return sum, avg
}

// 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
// 但是为了程序可读性，不建议如此使用
func test3(x, y int) (sum int, avg int) {
	sum = x + y
	// avg = sum / 2
	return
}

// 多个返回值的函数可以直接被其他函数调用
func test4(x, y int) int {
	return x + y
}

// 命名返回值会被本函数内生命的局部变量覆盖
func test5(x, y int) (z int) {
	z := x + y
	return
}