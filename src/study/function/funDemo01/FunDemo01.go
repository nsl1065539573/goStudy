package main

import "fmt"

func main(){
	s1 := test(func(x int) string {
		return string(x)
	}, 10)
	s2 := format(func(s string, x, y int) string {
        return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s1)
	fmt.Println(s2)
}

func test(fn func(x int) string, i int) string {
	return fn(i)
}

type Format func(s string, x, y int) string

func format(f Format, s string, x, y int) string {
	return f(s, x, y)
}

func sum(x, y int) int {
	return x + y
}