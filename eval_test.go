package main

import (
	"testing"

	"github.com/jiro4989/xlsxlang/token"
	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		desc   string
		tokens []token.Token
		want   token.Token
	}{
		{
			desc: "正常系: 加算",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("+"),
						token.NewIntToken(1),
						token.NewIntToken(2),
					},
				},
			},
			want: token.NewIntToken(3),
		},
		{
			desc: "正常系: 減算",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("-"),
						token.NewIntToken(2),
						token.NewIntToken(1),
					},
				},
			},
			want: token.NewIntToken(1),
		},
		{
			desc: "正常系: 乗算",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("*"),
						token.NewIntToken(3),
						token.NewIntToken(5),
					},
				},
			},
			want: token.NewIntToken(15),
		},
		{
			desc: "正常系: 除算",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("/"),
						token.NewIntToken(10),
						token.NewIntToken(5),
					},
				},
			},
			want: token.NewIntToken(2),
		},
		{
			desc: "正常系: 余剰",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("%"),
						token.NewIntToken(10),
						token.NewIntToken(3),
					},
				},
			},
			want: token.NewIntToken(1),
		},
		{
			desc: "正常系: 累乗",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("**"),
						token.NewIntToken(2),
						token.NewIntToken(10),
					},
				},
			},
			want: token.NewIntToken(1024),
		},
		{
			desc: "正常系: atom bool",
			tokens: []token.Token{
				token.NewBoolToken(true),
			},
			want: token.NewBoolToken(true),
		},
		{
			desc: "正常系: atom int",
			tokens: []token.Token{
				token.NewIntToken(2),
			},
			want: token.NewIntToken(2),
		},
		{
			desc: "正常系: atom str",
			tokens: []token.Token{
				token.NewStrToken("hello"),
			},
			want: token.NewStrToken("hello"),
		},
		{
			desc: "正常系: atom nil",
			tokens: []token.Token{
				token.NewNilToken(),
			},
			want: token.NewNilToken(),
		},
		{
			desc: "正常系: ネストした演算, (+ (+ 1 2) (+ 3 4))",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("+"),
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								token.NewSymbolToken("+"),
								token.NewIntToken(1),
								token.NewIntToken(2),
							},
						},
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								token.NewSymbolToken("+"),
								token.NewIntToken(3),
								token.NewIntToken(4),
							},
						},
					},
				},
			},
			want: token.NewIntToken(10),
		},
		{
			desc: "正常系: ネストした演算, (+ (+ 1 (- 3 1)) (+ (+ 2 1) 4))",
			tokens: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("+"),
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								token.NewSymbolToken("+"),
								token.NewIntToken(1),
								{
									Kind: token.KindList,
									ValueList: []token.Token{
										token.NewSymbolToken("-"),
										token.NewIntToken(3),
										token.NewIntToken(1),
									},
								},
							},
						},
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								token.NewSymbolToken("+"),
								{
									Kind: token.KindList,
									ValueList: []token.Token{
										token.NewSymbolToken("+"),
										token.NewIntToken(2),
										token.NewIntToken(1),
									},
								},
								token.NewIntToken(4),
							},
						},
					},
				},
			},
			want: token.NewIntToken(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := Evaluate(tt.tokens)
			assert.Equal(tt.want, got)
		})
	}
}
