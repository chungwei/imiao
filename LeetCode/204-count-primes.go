package main

import (
	"fmt"
	"math"
)

/**
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/
func main() {
	fmt.Println(countPrimes(10)) // 4
	fmt.Println(countPrimes(5))  // 2
}

func countPrimes(n int) int {
	cnt := 0
	for i := 2; i < n; i++ {
		f := false
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				f = true
				break
			}
		}
		if !f {
			cnt++
		}
	}

	return cnt
}
