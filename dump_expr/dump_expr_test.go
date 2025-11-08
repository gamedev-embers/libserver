package dump_expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		expr := "(a + b + c) = (b + a + c)"
		tokens, err := tokenizer(expr)
		if err != nil {
			t.Fatalf("tokenizer failed: %v", err)
		}
		chars := []rune(expr)
		assert.Equal(t, len(chars), len(tokens))
		for i, c := range chars {
			assert.Equal(t, string(c), tokens[i].Value)
		}
	})
}

func TestDump(t *testing.T) {
	_cases := []struct {
		expr string
		args map[string]any
		want string
	}{
		{
			expr: "x + y - z",
			args: map[string]any{
				"x": 10,
				"y": 20,
				"z": 5,
			},
			want: "x[10] + y[20] - z[5]",
		},
		{
			expr: "value1 * value2 / value3",
			args: map[string]any{
				"value1": 3.14159,
				"value2": 2.71828,
				"value3": 1.61803,
			},
			want: "value1[3.1416] * value2[2.7183] / value3[1.6180]",
		},
	}
	for _, c := range _cases {
		t.Run(c.expr, func(t *testing.T) {
			got, err := Dump(c.expr, c.args)
			assert.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
