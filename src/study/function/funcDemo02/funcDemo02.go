package main

import "fmt"

func main(){
	fmt.Println(test(1, 2, 3, 4, 5))
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(test(arr...))
}

func test(args ...int) string {
	sum := 0
	for _, x := range args {
		sum = sum + x
	}
	return fmt.Sprintf("sum is %v", sum)
}
