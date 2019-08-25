package main

import (
	"fmt"
	"strconv"
)

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21

*/
func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(100))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1563847412), 1<<31-1)
}

func reverse(x int) int {
	flag := true
	xx := strconv.Itoa(x)
	if x < 0 {
		flag = false
		xx = xx[1:]
	}

	tmp := []byte(xx)
	ret := ``
	for i := len(tmp) - 1; i >= 0; i-- {
		ret += string(tmp[i])
	}

	t, err := strconv.Atoi(ret)
	if err != nil || t > (1<<32-1) {
		return 0
	}

	if !flag {
		return -t
	}
	return t

}
