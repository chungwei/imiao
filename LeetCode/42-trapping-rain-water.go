package main

import "fmt"

/**
42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	n := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//fmt.Println(trap(n))
	fmt.Println(trap1(n))
}

/**
从两端往中间逼近，记录左右两端高度最高值，那么对于这两端最高值中间部分，如果高度低于两端最高值，能接的雨水取决于两端最高值中的最小值。
*/
//func trap(height []int) int {
//	left, right := 0, len(height)-1
//	lt, rt, sum := 0, 0, 0
//
//	for left < right {
//		if height[left] < height[right] {
//			if lt <= height[left] {
//				lt = height[left]
//			} else {
//				sum += lt - height[left]
//			}
//			left++
//		} else {
//			if rt <= height[right] {
//				rt = height[right]
//			} else {
//				sum += rt - height[right]
//			}
//			right--
//		}
//	}
//	return sum
//}

/**
这种方法相对而言最好理解了。
先找到最高的那个柱子，把数组分成两部分，
对于两部分都已经确定了一个边界高度了，所以对剩余的每个柱子至于确定一边的边界高度值，就可以直接计算出能接的雨水了
*/
//func trap1(height []int) int {
//	right := len(height) - 1
//	max, lt, rt, sum := 0, 0, 0, 0
//
//	for k, v := range height {
//		if height[max] < v {
//			max = k
//		}
//	}
//
//	for i := 0; i < max; i++ {
//		if lt <= height[i] {
//			lt = height[i]
//		} else {
//			sum += lt - height[i]
//		}
//	}
//
//	for i := right; i > max; i-- {
//		if rt <= height[i] {
//			rt = height[i]
//		} else {
//			sum += rt - height[i]
//		}
//	}
//	return sum
//}

func trap1(height []int) int {
	right := len(height) - 1
	if right < 2 { // 说明少于3个元素 直接返回0
		return 0
	}

	max, idx, sum := height[0], 0, 0
	for k, v := range height { // 找到最大值
		if max < v {
			max = v
			idx = k
		}
	}

	tmp := 0
	for i := 0; i < idx; i++ {
		if tmp <= height[i] {
			tmp = height[i]
		} else {
			sum += tmp - height[i]
		}
	}

	tmp = 0
	for i := right; i > idx; i-- {
		if tmp <= height[i] {
			tmp = height[i]
		} else {
			sum += tmp - height[i]
		}
	}
	return sum
}
