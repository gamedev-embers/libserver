package dump_expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	tokens, err := tokenizer("a + b - c - a")
	if err != nil {
		t.Fatalf("tokenizer failed: %v", err)
	}
	assert.Equal(t, 13, len(tokens))
	assert.Equal(t, "a", tokens[0].Value)
	assert.Equal(t, "b", tokens[4].Value)
	assert.Equal(t, "c", tokens[8].Value)
	assert.Equal(t, "a", tokens[12].Value)
}

func TestDump(t *testing.T) {
	s, err := Dump("a + b - c + a", map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
	})
	if err != nil {
		t.Fatalf("tokenizer failed: %v", err)
	}
	assert.Equal(t, "a[1] + b[2] - c[3] + a[1]", s)
}
