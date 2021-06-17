package iteration

import "strings"

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

func IsSameString(a string, b string) bool {
	var isSame = strings.Compare(a, b)
	if isSame == 0 {
		return true
	} else {
		return false
	}

}
