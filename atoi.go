package piscine

func Atoi(s string) int {
	var res int
	bang := "positive"
	if s == "" {
		res = 0
	} else {
		if s[0] == '-' {
			bang = "negative"
			s = s[1:]
		} else if s[0] == '+' {
			s = s[1:]
		}
		if len(s) > 0 {
			for i := range s {
				if s[i] < 48 || s[i] > 57 {
					res = 0
					break
				}
				res = res*10 + int(s[i]-'0')
			}
			if bang == "negative" {
				res *= -1
			}

		}

	}
	return res
}
