package db

import (
	"fmt"
	"strings"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func main() {
	fmt.Println(lengthOfLongestSubstring1(`abcabcbb`))
	fmt.Println(lengthOfLongestSubstring(`abcabcbb`))
	/*fmt.Println(lengthOfLongestSubstring1(`bbbbb`))
	fmt.Println(lengthOfLongestSubstring1(`pwwkew`))
	fmt.Println(lengthOfLongestSubstring1(` `))
	fmt.Println(lengthOfLongestSubstring1(`a`))
	fmt.Println(lengthOfLongestSubstring1(`au`))
	fmt.Println(lengthOfLongestSubstring1(`dvdf`))
	fmt.Println(lengthOfLongestSubstring1("asjrgapa"))
	fmt.Println(lengthOfLongestSubstring1("aab"))*/
}

func lengthOfLongestSubstring1(s string) int {
	length := len(s)
	ret, a := 0, ``
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			t := string(s[j])
			if strings.Contains(a, t) {
				if ret < len(a) {
					ret = len(a)
				}
				//fmt.Println(a, ret)
				a = ``
				break
			}
			a += t
		}
	}
	if 0 == ret {
		return length
	}
	return ret
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	cur := 0
	ints := make(map[int32]int)
	for k, v := range s {
		key, ok := ints[v]
		fmt.Println(k, key, ok)

		dis := k - key
		if !ok || dis > cur {
			cur++
		} else {
			cur = dis
		}
		ints[v] = k
		if max < cur {
			max = cur
		}
	}
	return max
}
