package main

import "fmt"

func  main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
    // 思路 滑动窗口 窗口大小为needle的长度，当窗口右边界达到haystack的右端边界后，结束循环
	var needlelen int = len(needle)
	fmt.Printf("needlelen: %v \n", needlelen)
	if (needlelen == 0){
		return 0
	}
    var hayLen int = len(haystack)
    if (needlelen > hayLen) {
        return -1
    }
    var needArr []byte = []byte(needle)
    var haystackArr []byte = []byte(haystack)
    var left, right int;
    left = 0;
	right = needlelen - 1
	fmt.Println(needArr)
    for ; right < hayLen; right++ {
		for i := left; i <= right; i++ {
			fmt.Printf("i: %v, left: %v, right: %v\n", i, left, right)
			fmt.Printf("left: %v, rigth: %v \n",haystackArr[left], haystackArr[right])
			if needArr[i - left] != haystackArr[i]{
				break
			}
			if i == right  {
				return left
			}
		}
		fmt.Printf("进入第%v次循环\n", left)
		left += 1
    }
    return -1
}