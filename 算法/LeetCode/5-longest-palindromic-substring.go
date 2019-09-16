package db

import (
	"fmt"
)

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
func main() {
	// 判断回文
	fmt.Println(isPalindrome(`ss`))
	fmt.Println(isPalindrome(`s`))
	fmt.Println(isPalindrome(``))
	fmt.Println(isPalindrome(`abc`))
	fmt.Println(isPalindrome(`aba`))

	fmt.Println(longestPalindrome(`babad`))
	fmt.Println(longestPalindrome(`abad`))
	fmt.Println(longestPalindrome(`cc`))
	fmt.Println(longestPalindrome(`aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`))
	fmt.Println(longestPalindrome(`cccc`))
	fmt.Println(longestPalindrome(`cbbd`))
	fmt.Println(longestPalindrome(`abcda`))
	fmt.Println(longestPalindrome(`abcdr`))
}

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	max := 0
	ret := ``
	for i := 0; i < length-1; i++ {
		for j := length; j > i+1; j-- {
			tmp := s[i:j]
			//fmt.Println(`==`, tmp)
			if isPalindrome(tmp) && len(tmp) > max {
				ret = tmp
				max = len(tmp)
			}
		}
	}
	if ret == `` {
		return s[0:1]
	}
	return ret

}

func isPalindrome(s string) bool {
	// 判断回文
	length := len(s)
	if length < 2 {
		return true
	}

	j := length - 1
	for i := 0; i < length/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
