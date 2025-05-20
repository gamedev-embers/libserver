package textutils

import "unicode"

const (
	ZWJ  = "\u200D"
	ZWNJ = "\u200C"
)

func CheckText(chars []rune, min, max int) bool {
	if !IsLegal(chars) {
		return false
	}
	size := LengthOf(chars)
	return size >= min && size <= max
}

// 视觉长度，中文算2个
func LengthOf(chars []rune) int {
	l := 0
	for _, ch := range chars {
		if ch >= 0x80 {
			l += 2
		} else {
			l++
		}
	}
	return l
}

// 只允许字母、数字、中文
func IsLegal(chars []rune) bool {
	for _, c := range chars {
		if !unicode.IsPrint(c) {
			// 不可见字符
			return false
		}

		if unicode.IsControl(c) {
			// 控制字符
			return false
		}

		if unicode.IsSpace(c) {
			// 空格
			return false
		}

		if unicode.IsPunct(c) {
			// 标点
			return false
		}

		if unicode.IsSymbol(c) {
			// 符号
			return false
		}
		if unicode.IsMark(c) {
			// 商标字符
			return false
		}
	}
	return true
}
