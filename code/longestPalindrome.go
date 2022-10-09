package code

// 最长不重复子串
func LongestPalindrome(s string) string {
	n := 0
	r := ""
	l := len(s)

	if l == 1 || (l == 2 && s[0] == s[1]) {
		return s
	}
	if l == 2 {
		return string(s[0])
	}

	for i := range s {

		for j := 0; j < l; j++ {

			if i-j < 0 {
				break
			}

			if i+j > l-1 {
				break
			}

			if s[i-j] == s[i+j] {
				if n < len(s[i-j:i+j+1]) {
					n = len(s[i-j : i+j+1])
					r = s[i-j : i+j+1]
				}
			} else if s[i] == s[i+1] && s[i-j] == s[i+1+j] {
				if s[i-j] == s[i+1+j] {
					if n < len(s[i-j:i+1+j+1]) {
						n = len(s[i-j : i+1+j+1])
						r = s[i-j : i+1+j+1]
					}
				}
			} else {
				break
			}
		}
	}
	return r
}
