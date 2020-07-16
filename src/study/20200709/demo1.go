package main

import "fmt"

func main(){
	var a []int = make([]int, 3)
	a = append(a, 10)
	fmt.Println(a)
}