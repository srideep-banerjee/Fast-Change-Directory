package validation

import "strings"

var allowed = "abcdefghijklmnopqrstuvwxyz"

func IsTagValid(tag string) (bool, rune) {
    for _, ch := range tag {
        if !isCharAllowed(ch) {
            return false, ch
        }
    }
    return true, '0'
}

func isCharAllowed(ch rune) bool {
	if strings.ContainsRune(allowed, ch) {
		return true
	}
	if strings.ContainsRune(strings.ToUpper(allowed), ch) {
		return true
	}
	if '0' <= ch && ch <= '9' {
		return true
	}
	return false
}
