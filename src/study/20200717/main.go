
package main

import (
	"fmt"
	"strings"
)

func main() {
    // s := "helloworld"
    // fmt.Println(s[0:2])//he,右半包含
    // fmt.Println(s[2:])//lloworld
    // fmt.Println(s[:9])//hellow
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}


func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    str := strs[0]
    for i := 1; i < len(strs); i++ {
        for {
            if len(str) == 0 {
                return ""
            }
            if !strings.HasPrefix(strs[i], str) {
                str = str[: len(str) - 1]
            } else {
				break
			}
        }
    }
    return str
}