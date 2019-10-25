package string

func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < ls; i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
