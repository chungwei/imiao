package main

import "fmt"

func main() {
	n := []int{1, 0, 2}
	fmt.Println(candy(n))

	n = []int{1, 2, 2}
	fmt.Println(candy(n))
}

func candy(ratings []int) int {
	c := make([]int, len(ratings))
	//for i := 0; i < len(ratings); i++ {
	//c[i] = 1
	//}
	fmt.Println(c)

	//for i := 0; i < len(ratings)-1; i++ {
	//if c[i] < c[i+1] {
	//	c[i+1] = c[i] + 1
	//}
	//}
	//fmt.Println(c)

	return sum
}
