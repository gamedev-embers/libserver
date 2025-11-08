package dump_expr

import (
	"fmt"
	"strings"
)

const (
	TokenType_Unknown = iota
	TokenType_Space
	TokenType_Newline
	TokenType_Identifier
	TokenType_Operator
	TokenType_Number
)

type Token struct {
	Type  int
	Value string
	Start int
	End   int
	Index int
}

func (t Token) String() string {
	return fmt.Sprintf("%s<%d>:%d~%d", t.Value, t.Type, t.Start, t.End)
}

type Var struct {
	Name  string
	Token Token
}

func tokenizer(f string) ([]*Token, error) {
	state := 0 // 0-start
	var cur *Token

	tokens := []*Token{}
	chars := []rune(f)
	tokenStart := func(offset int, c string, tp int) *Token {
		if cur == nil {
			cur = &Token{}
		}
		cur.Value += c
		cur.Start = offset
		cur.End = offset + 1
		cur.Type = tp
		return cur
	}

	tokenContinue := func(offset int, c string) {
		if cur == nil {
			cur = &Token{}
		}
		cur.Value += c
		cur.End = offset + 1
	}
	tokenEnd := func() {
		if cur == nil {
			return
		}
		tokens = append(tokens, cur)
		cur = nil
		state = 0
	}

	for i := 0; i < len(chars); i++ {
		c := chars[i]

		switch {
		case c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_':
			switch state {
			case 0:
				state = 1
				tokenStart(i, string(c), TokenType_Identifier)
			case 1:
				if cur.Type != TokenType_Identifier {
					tokenEnd()
					tokenStart(i, string(c), TokenType_Identifier)
				} else if cur.Type == TokenType_Identifier {
					tokenContinue(i, string(c))
				}
			}
		case c >= '0' && c <= '9':
			switch state {
			case 0:
				tokenStart(i, string(c), TokenType_Number)
			case 1:
				tokenContinue(i, string(c))
			}
		case c == '+' || c == '-' || c == '*' || c == '/' || c == '%' || c == '=':
			switch state {
			case 0:
				tokenStart(i, string(c), TokenType_Operator)
				tokenEnd()
			case 1:
				state = 0
				tokenEnd()
				tokenStart(i, string(c), TokenType_Operator)
				tokenEnd()
			}
		case c == ' ' || c == '\t':
			switch state {
			case 0:
				tokenStart(i, string(c), TokenType_Space)
				tokenEnd()
			case 1:
				state = 0
				tokenEnd()
				tokenStart(i, string(c), TokenType_Space)
				tokenEnd()
			}
		case c == '\n' || c == '\r':
			switch state {
			case 0:
				tokenStart(i, string(c), TokenType_Newline)
				tokenEnd()
			case 1:
				state = 0
				tokenEnd()
				tokenStart(i, string(c), TokenType_Newline)
				tokenEnd()
			}
		case c == '(' || c == ')' || c == ',' || c == '.' || c == '[' || c == ']':
			switch state {
			case 0:
				tokenStart(i, string(c), TokenType_Newline)
				tokenEnd()
			case 1:
				state = 0
				tokenEnd()
				tokenStart(i, string(c), TokenType_Newline)
				tokenEnd()
			}
		default:
			switch state {
			case 0:
				state = 1
				tokenStart(i, string(c), TokenType_Identifier)
			case 1:
				tokenContinue(i, string(c))
			}
		}
	}
	tokenEnd()
	return tokens, nil
}

func Dump(f string, args map[string]any) (string, error) {
	return DumpWith(f, args)
}

func DumpWith(f string, args map[string]any, opts ...Option) (string, error) {
	_opts := buildOptions(opts...)

	tokens, err := tokenizer(f)
	if err != nil {
		return "", err
	}
	var texts = []string{}
	var vars = []*Token{}
	var sb *strings.Builder
	onString := func(s string) {
		if sb == nil {
			sb = &strings.Builder{}
		}
		sb.WriteString(s)
	}
	onVar := func(tk *Token) {
		vars = append(vars, tk)
		tk.Index = len(texts)
		texts = append(texts, tk.Value)
	}
	onStringEnd := func() {
		if sb == nil {
			return
		}
		texts = append(texts, sb.String())
		sb = nil
	}

	for i, tk := range tokens {
		switch tk.Type {
		case TokenType_Space:
			onString(tk.Value)
		case TokenType_Identifier:
			onStringEnd()
			onVar(tk)
		case TokenType_Operator:
			onString(tk.Value)
		case TokenType_Number:
			onString(tk.Value)
		case TokenType_Newline:
			onString(tk.Value)
		default:
			panic(fmt.Errorf("invalid token. i:%d token:%v", i, tk))
		}
	}
	onStringEnd()
	for _, _var := range vars {
		if v, exists := args[_var.Value]; exists {
			// old := texts[_var.Index]
			texts[_var.Index] = _opts.FormatVar(_var.Value, v)
		}
	}
	return strings.Join(texts, ""), nil
}
