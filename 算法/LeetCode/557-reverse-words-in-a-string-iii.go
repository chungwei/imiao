package main

import (
	"fmt"
	"strings"
)

/**
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:
输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
*/
func main() {
	x := "Let's take LeetCode contest"
	fmt.Println(reverseWords(x))

}
func reverseWords(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	tmp := strings.Split(s, " ")

	for i := 0; i < len(tmp); i++ {
		tmp[i] = swap(tmp[i])
	}
	return strings.Join(tmp, " ")
}

func swap(s string) string {
	ss := []byte(s)
	l := len(s)
	j := l - 1
	for i := 0; i < l/2; i++ {
		// 两两交换位置 这里要注意数组长度是奇数和偶数的差异
		t := ss[i]
		ss[i] = ss[j]
		ss[j] = t
		j--
	}
	return string(ss)
}
