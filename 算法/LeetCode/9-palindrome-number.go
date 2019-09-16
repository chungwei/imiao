package db

import (
	"fmt"
	"strconv"
)

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/
func main() {
	fmt.Println(isPalindrome1(123))
	fmt.Println(isPalindrome1(-123))
	fmt.Println(isPalindrome1(100))
	fmt.Println(isPalindrome1(11))
}

func isPalindrome1(x int) bool {
	s := strconv.Itoa(x)
	if len(s) == 1 {
		return true
	}
	if x < 0 {
		return false
	}
	length := len(s)
	j := length - 1
	bb := []byte(s)
	for i := 0; i < length/2; i++ {
		if bb[i] != bb[j] {
			return false
		}
		j--
	}
	return true
}
