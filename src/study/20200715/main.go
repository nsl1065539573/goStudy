package main

import "fmt"

func main(){
	fmt.Println(countAndSay(5))
}

type ListNode struct {
     Val int
     Next *ListNode
}

// 删除倒数第K个节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 求出链表长度
    temp := head
    length := 0
    for {
        if temp.Next == nil {
            break
        }
        temp = temp.Next
        length += 1
	}
	res := head
    for i := 0; i < length - 1 - n; i++ {
        res = res.Next
    }
    return res.Next
}

// 思路： 递归  当n == 1 时  return 1
// 当n > 1 时， return say(n)
func countAndSay(n int) string {
	// fmt.Printf("第%v次调用\n", n)
    if n == 1 {
        return "1"
    }
    return say(countAndSay(n - 1))
}

func say(str string) string {
	fmt.Printf("调用一次say，str: %v\n", str)
	var strArr []byte = []byte(str)
	fmt.Println(strArr)
    var item, times byte
    item = strArr[0]
    times = 48
    var res []byte = make([]byte, 0)
    for i := 0; i < len(strArr); i++ {
        if item != strArr[i] {
            res = append(res, times)
            res = append(res, item)
            item = strArr[i]
            times = 49
        } else {
            times += 1
        }
    }
    res = append(res, times)
	res = append(res, item)
	return string(res[:])
}