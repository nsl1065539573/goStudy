package main

import "fmt"


func main(){
	i := 10
	arr := make([]int, 0)
	changeArr(i, &arr)
	fmt.Println(arr)
}

func changeArr(i int, arr *[]int) {
	if i == 0 {
		return
	}
	j := i
	*arr = append(*arr, j)
	fmt.Printf("打印i:%p\n", arr)
	changeArr(i - 1, arr)
}