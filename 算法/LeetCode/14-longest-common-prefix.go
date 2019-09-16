package db

import "fmt"

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

*/
func main() {
	s := []string{`flww`, `flvvvvwww`, `fl22owww`}
	fmt.Println(longestCommonPrefix(s))

	s = []string{`xflww`, `flvvvvwww`, `fl22owww`}
	fmt.Println(longestCommonPrefix(s))

	s = []string{``}
	fmt.Println(longestCommonPrefix(s))

	s = []string{}
	fmt.Println(longestCommonPrefix(s))

	s = []string{`aca`, `cba`}
	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ``
	}
	minLen, str := len(strs[0]), []byte(strs[0])
	for _, v := range strs {
		if l := len(v); minLen > l {
			minLen = l
			str = []byte(v)
		}
	}
	if minLen == 0 {
		return ``
	}

	ret := ``
	for i := 0; i < minLen; i++ {
		match := true
		for j := 0; j < len(strs); j++ {
			bb := []byte(strs[j])
			//fmt.Println(string(str[i]), string(bb[i]), match)
			if bb[i] != str[i] {
				match = false
				return ret
			}
		}
		if match {
			ret += string(str[i])
		}
	}
	return ret

}
