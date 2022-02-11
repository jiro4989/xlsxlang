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
			desc: "正常系: シンプルな足し算",
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
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := Evaluate(tt.tokens)
			assert.Equal(tt.want, got)
		})
	}
}
