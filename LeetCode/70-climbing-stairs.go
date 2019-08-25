package main

import "fmt"

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
*/
func main() {
	fmt.Println(climbStairs1(2))  // 2
	fmt.Println(climbStairs1(3))  // 3
	fmt.Println(climbStairs1(4))  // 5
	fmt.Println(climbStairs1(5))  // 8
	fmt.Println(climbStairs1(44)) // 1134903170
	fmt.Println(climbStairs2(2))  // 2
	fmt.Println(climbStairs2(3))  // 3
	fmt.Println(climbStairs2(4))  // 5
	fmt.Println(climbStairs2(5))  // 8
	fmt.Println(climbStairs2(6))  // 13
	fmt.Println(climbStairs2(44)) // 1134903170
}

func climbStairs1(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

func climbStairs2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	m := make(map[int]int, n+1)
	m[0], m[1], m[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		m[i] = m[i-2] + m[i-1]
	}
	return m[n]
}
