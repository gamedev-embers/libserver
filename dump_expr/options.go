package dump_expr

import "fmt"

type options struct {
	FormatVar func(name string, value any) string
}

type Option func(*options)

func buildOptions(opts ...Option) *options {
	_opts := defaultOptions()
	for _, o := range opts {
		o(_opts)
	}
	return _opts
}

func defaultOptions() *options {
	return &options{
		FormatVar: func(name string, value any) string {
			switch value.(type) {
			case float32, float64:
				return fmt.Sprintf("%s[%.4f]", name, value)
			default:
				return fmt.Sprintf("%s[%v]", name, value)
			}
		},
	}
}

func WithFormatVar(f func(name string, value any) string) Option {
	return func(o *options) {
		o.FormatVar = f
	}
}
