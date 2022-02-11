package main

import (
	"testing"

	"github.com/jiro4989/xlsxlang/token"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc    string
		program string
		want    []token.Token
		wantErr bool
	}{
		{
			desc:    "正常系: trueのみ",
			program: "true",
			want: []token.Token{
				token.NewBoolToken(true),
			},
		},
		{
			desc:    "正常系: falseのみ",
			program: "false",
			want: []token.Token{
				token.NewBoolToken(false),
			},
		},
		{
			desc:    "正常系: intのみ",
			program: "1",
			want: []token.Token{
				token.NewIntToken(1),
			},
		},
		{
			desc:    "正常系: intのみ",
			program: "255",
			want: []token.Token{
				token.NewIntToken(255),
			},
		},
		{
			desc:    "正常系: stringのみ",
			program: `"hello"`,
			want: []token.Token{
				token.NewStrToken("hello"),
			},
		},
		{
			desc:    "正常系: nilのみ",
			program: `nil`,
			want: []token.Token{
				token.NewNilToken(),
			},
		},
		{
			desc:    "正常系: symbolのみ",
			program: "exists?",
			want: []token.Token{
				token.NewSymbolToken("exists?"),
			},
		},
		{
			desc:    "正常系: symbolのみ",
			program: "+",
			want: []token.Token{
				token.NewSymbolToken("+"),
			},
		},
		{
			desc:    "正常系: 複数のatom",
			program: `+ 1 true "hello"`,
			want: []token.Token{
				token.NewSymbolToken("+"),
				token.NewIntToken(1),
				token.NewBoolToken(true),
				token.NewStrToken("hello"),
			},
		},
		{
			desc:    "正常系: 複数のatom (改行)",
			program: "+\n1\ntrue\n\"hello\"",
			want: []token.Token{
				token.NewSymbolToken("+"),
				token.NewIntToken(1),
				token.NewBoolToken(true),
				token.NewStrToken("hello"),
			},
		},
		{
			desc:    "正常系: 複数のlist",
			program: `(+ 1 2) (+ 3 4)`,
			want: []token.Token{
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
		{
			desc:    "正常系: 要素が1つのみのリスト",
			program: `(sym)`,
			want: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("sym"),
					},
				},
			},
		},
		{
			desc:    "正常系: 単純な1つのリスト",
			program: `(hello 1 "world" nil true)`,
			want: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("hello"),
						token.NewIntToken(1),
						token.NewStrToken("world"),
						token.NewNilToken(),
						token.NewBoolToken(true),
					},
				},
			},
		},
		{
			desc:    "正常系: ネストしたリスト",
			program: `(hello (foo 1))`,
			want: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("hello"),
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								token.NewSymbolToken("foo"),
								token.NewIntToken(1),
							},
						},
					},
				},
			},
		},
		{
			desc:    "正常系: ネストしたリスト 2",
			program: `(= (+ 1 2) (- 3 2))`,
			want: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("="),
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
								token.NewSymbolToken("-"),
								token.NewIntToken(3),
								token.NewIntToken(2),
							},
						},
					},
				},
			},
		},
		{
			desc:    "正常系: ネストしまくり",
			program: `(+ 1 (- (+ 1 1) (* 3 (/ 4 5))))`,
			want: []token.Token{
				{
					Kind: token.KindList,
					ValueList: []token.Token{
						token.NewSymbolToken("+"),
						token.NewIntToken(1),
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								token.NewSymbolToken("-"),
								{
									Kind: token.KindList,
									ValueList: []token.Token{
										token.NewSymbolToken("+"),
										token.NewIntToken(1),
										token.NewIntToken(1),
									},
								},
								{
									Kind: token.KindList,
									ValueList: []token.Token{
										token.NewSymbolToken("*"),
										token.NewIntToken(3),
										{
											Kind: token.KindList,
											ValueList: []token.Token{
												token.NewSymbolToken("/"),
												token.NewIntToken(4),
												token.NewIntToken(5),
											},
										},
									},
								},
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

			got := parser.GetTokens()
			assert.Equal(tt.want, got)
		})
	}
}
