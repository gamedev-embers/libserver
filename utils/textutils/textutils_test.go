package textutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func T(s string) []rune {
	return []rune(s)
}

func T2(s string) [][]rune {
	lst := [][]rune{}
	for _, c := range s {
		lst = append(lst, []rune{c})
	}
	return lst
}

func Test_IsLegal(t *testing.T) {
	t.Run("no symbals", func(t *testing.T) {
		assert.False(t, IsLegal(T("Invalid Name!")))
		assert.False(t, IsLegal(T("Invalid @!")))
		assert.False(t, IsLegal(T("[]")))
		assert.False(t, IsLegal(T("【")))
		assert.False(t, IsLegal(T("】")))
		for _, c := range T2("!@#$%^&*()_+{}[]:\"|';,./<>?【】") {
			assert.False(t, IsLegal(c))
		}
	})

	t.Run("no spaces", func(t *testing.T) {
		for _, c := range T2(" \t\n\r") {
			assert.False(t, IsLegal(c))
		}
	})

	t.Run("no emoji", func(t *testing.T) {
		for _, c := range T2("⌚️💖💔💡💣💥💦💨💩💪💫💬🕵️‍♂️‍♀️🕶️🕷️🕸️🕹️🖤") {
			assert.False(t, IsLegal(c), "表情符号: %s %v", string(c), c)
		}
	})

	t.Run("no ZWJ/ZWNJ", func(t *testing.T) {
		c := T("\u200D")
		assert.Equal(t, []rune{8205}, c, "ZWJ: %s %v", string(c), c)
		assert.False(t, IsLegal(c), "ZWJ: %s %v", string(c), c)

		c = T("\u200C")
		assert.Equal(t, []rune{8204}, c, "ZWNJ: %s %v", string(c), c)
		assert.False(t, IsLegal(c), "ZWNJ: %s %v", string(c), c)
	})

	t.Run("allow numbers", func(t *testing.T) {
		for _, c := range T2("0123456789") {
			assert.True(t, IsLegal(c), "%s %v", string(c), c)
		}
	})

	t.Run("allow letters", func(t *testing.T) {
		for _, c := range T2("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			assert.True(t, IsLegal(c), "%s %v", string(c), c)
		}
	})

	t.Run("allow CJK", func(t *testing.T) {
		assert.True(t, IsLegal(T("中文")))
		assert.True(t, IsLegal(T("中文测试")))
		assert.True(t, IsLegal(T("中文测试123")))
		assert.False(t, IsLegal(T("中文测试123!@#$%^&*()_+{}[]:\"|';,./<>?")))
	})

}

func Test_LengthOf(t *testing.T) {
	assert.Equal(t, 10, LengthOf(T("你好啊你好")))
	assert.Equal(t, 9, LengthOf(T("你好啊你1")))
	assert.Equal(t, 10, LengthOf(T("你好啊你1a")))
	assert.Equal(t, 10, LengthOf(T("1a你好啊你")))
	assert.Equal(t, 10, LengthOf(T("abcdefghij")))
	assert.Equal(t, 10, LengthOf(T("abcdefghij")))
}

func Test_CheckName(t *testing.T) {
	t.Run("letters", func(t *testing.T) {
		assert.True(t, CheckText(T("abcdefghij"), 1, 10))
		assert.False(t, CheckText(T("abcdefghijl"), 1, 10))
	})

	t.Run("numbers", func(t *testing.T) {
		assert.True(t, CheckText(T("1234567890"), 1, 10))
		assert.False(t, CheckText(T("12345678901"), 1, 10))
	})

	t.Run("CJK", func(t *testing.T) {
		assert.True(t, CheckText(T("中文测试"), 1, 10))
		assert.True(t, CheckText(T("中文测试12"), 1, 10))
		assert.False(t, CheckText(T("中文测试123"), 1, 10))
	})
}
