package week_0

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	gcd := ""
	len1 := len(str1)
	len2 := len(str2)
	for i := 1; i <= len1; i++ {
		if i > len2 {

			break
		}
		if len1%i == 0 && len2%i == 0 {
			divisor := str1[0:i]
			if strings.Repeat(divisor, len1/i) == str1 && strings.Repeat(divisor, len2/i) == str2 {
				gcd = divisor
			}
		}
	}
	return gcd
}
