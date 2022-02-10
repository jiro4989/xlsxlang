package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc    string
		program string
		want    Token
		wantErr bool
	}{
		{
			desc:    "正常系: trueのみ",
			program: "true",
			want: Token{
				Kind:      kindBool,
				ValueBool: true,
			},
		},
		{
			desc:    "正常系: falseのみ",
			program: "false",
			want: Token{
				Kind:      kindBool,
				ValueBool: false,
			},
		},
		{
			desc:    "正常系: intのみ",
			program: "1",
			want: Token{
				Kind:     kindInt,
				ValueInt: 1,
			},
		},
		{
			desc:    "正常系: intのみ",
			program: "255",
			want: Token{
				Kind:     kindInt,
				ValueInt: 255,
			},
		},
		{
			desc:    "正常系: stringのみ",
			program: `"hello"`,
			want: Token{
				Kind:     kindStr,
				ValueStr: "hello",
			},
		},
		{
			desc:    "正常系: symbolのみ",
			program: "exists?",
			want: Token{
				Kind:        kindSymbol,
				ValueSymbol: "exists?",
			},
		},
		{
			desc:    "正常系: symbolのみ",
			program: "+",
			want: Token{
				Kind:        kindSymbol,
				ValueSymbol: "+",
			},
		},
		{
			desc:    "正常系: 単純な1つのリスト",
			program: `(hello 1 "world" true)`,
			want: Token{
				Kind: kindList,
				ValueList: []Token{
					{
						Kind:        kindSymbol,
						ValueSymbol: "hello",
					},
					{
						Kind:     kindInt,
						ValueInt: 1,
					},
					{
						Kind:     kindStr,
						ValueStr: "world",
					},
					{
						Kind:      kindBool,
						ValueBool: true,
					},
				},
			},
		},
		{
			desc:    "正常系: ネストしたリスト",
			program: `(hello (foo 1) (bar true))`,
			want: Token{
				Kind: kindList,
				ValueList: []Token{
					{
						Kind:        kindSymbol,
						ValueSymbol: "hello",
					},
					{
						Kind: kindList,
						ValueList: []Token{
							{
								Kind:        kindSymbol,
								ValueSymbol: "foo",
							},
							{
								Kind:     kindInt,
								ValueInt: 1,
							},
						},
					},
					{
						Kind: kindList,
						ValueList: []Token{
							{
								Kind:        kindSymbol,
								ValueSymbol: "bar",
							},
							{
								Kind:      kindBool,
								ValueBool: true,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			parser, err := parse(tt.program)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(parser)
				return
			}

			got := parser.Eval.Token
			assert.Equal(tt.want, got)
		})
	}
}
