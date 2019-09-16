package main

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

输入: 1
输出: true
解释: 20 = 1
示例 2:

输入: 16
输出: true
解释: 24 = 16
示例 3:

输入: 218
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))

	fmt.Println(isPowerOfTwo1(1))
	fmt.Println(isPowerOfTwo1(16))
	fmt.Println(isPowerOfTwo1(4096 * 4096 * 4096 * 4096 * 4096))
}

func isPowerOfTwo(n int) bool {
	// 一个数如果是2的n次方，那么这个数二进制中只有一位是1，其余是0
	// 如果该数减1，则其二进制与上面的完全不同，
	// 利用这点，如果n是2的n次方，那么n&(n-1)必为0
	if n < 1 {
		return false
	}
	return (n & (n - 1)) == 0
}

func isPowerOfTwo1(n int) bool {
	if n < 1 {
		return false
	} else if n == 1 {
		return true // 2^0
	}

	// 暴力解决
	for n > 1 {
		if n%2 != 0 { // 不是2的倍数  肯定不是2的幂
			return false
		}
		n = n / 2
	}
	return true
}
