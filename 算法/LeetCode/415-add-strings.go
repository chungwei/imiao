package db

/**
415. 字符串相加

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n1, n2 := `1`, `9`
	fmt.Println(addStrings(n1, n2))

	n1, n2 = `0`, `0`
	fmt.Println(addStrings(n1, n2))

	n1, n2 = `584`, `18`
	fmt.Println(addStrings(n1, n2))

}

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	tmp := l1 - l2
	if tmp < 0 {
		tmp = -tmp
	}
	s := strings.Repeat(`0`, tmp) // 补零

	// 两个字符串长度相等了
	if l1 < l2 {
		num1 = s + num1
	} else {
		num2 = s + num2
	}

	lb1, lb2 := []byte(num1), []byte(num2)
	ret, c := ``, 0
	for i := len(num1) - 1; i >= 0; i-- {
		i1, _ := strconv.Atoi(string(lb1[i]))
		i2, _ := strconv.Atoi(string(lb2[i]))
		sum := i1 + i2 + c // 对应位相加
		ret = strconv.Itoa(sum%10) + ret
		c = sum / 10 // 进位
	}
	if c > 0 {
		ret = `1` + ret
	}
	return ret
}
