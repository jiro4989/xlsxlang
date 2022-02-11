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
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := Evaluate(tt.tokens)
			assert.Equal(tt.want, got)
		})
	}
}
