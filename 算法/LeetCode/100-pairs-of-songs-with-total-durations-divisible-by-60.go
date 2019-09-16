package db

import "fmt"

/**

 */
func main() {
	n1 := []int{30, 20, 150, 100, 40}
	fmt.Println(numPairsDivisibleBy60(n1))

	n1 = []int{60, 60, 60}
	fmt.Println(numPairsDivisibleBy60(n1))

	n1 = []int{60, 60, 60}
	fmt.Println(numPairsDivisibleBy601(n1))

	n1 = []int{30, 20, 150, 100, 40}
	fmt.Println(numPairsDivisibleBy601(n1))

	n1 = []int{29, 20, 151, 100, 40}
	fmt.Println(numPairsDivisibleBy601(n1))
}

func numPairsDivisibleBy60(time []int) int {
	t := 0
	for i := 0; i < len(time)-1; i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				t++
			}
		}
	}
	return t
}

func numPairsDivisibleBy601(time []int) int {
	// [30/30,20/40,150/30,100/20,40/20]
	m := make([]int, 60)
	cnt := 0
	for _, n := range time {
		if n%60 == 0 {
			cnt += m[0]
		} else {
			cnt += m[60-n%60]
		}
		m[n%60]++
		fmt.Println(m, 60-n%60, cnt)

	}
	return cnt
}
