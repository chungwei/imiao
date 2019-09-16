package db

import "fmt"

/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

*/
func main() {
	x := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(x))
	fmt.Println(maxArea1(x))
}

func maxArea(height []int) int {
	l, max := len(height), 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			t := height[i]
			if t > height[j] {
				t = height[j]
			}

			if a := (j - i) * t; max < a {
				max = a
			}
		}
	}

	return max
}
func maxArea1(height []int) int {
	i, j, max := 0, len(height)-1, 0

	for i < j {
		t := height[i]
		if t > height[j] {
			t = height[j]
		}

		if a := (j - i) * t; max < a {
			max = a
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return max
}
