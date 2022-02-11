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
				{
					Kind:      token.KindBool,
					ValueBool: true,
				},
			},
		},
		{
			desc:    "正常系: falseのみ",
			program: "false",
			want: []token.Token{
				{
					Kind:      token.KindBool,
					ValueBool: false,
				},
			},
		},
		{
			desc:    "正常系: intのみ",
			program: "1",
			want: []token.Token{
				{
					Kind:     token.KindInt,
					ValueInt: 1,
				},
			},
		},
		{
			desc:    "正常系: intのみ",
			program: "255",
			want: []token.Token{
				{
					Kind:     token.KindInt,
					ValueInt: 255,
				},
			},
		},
		{
			desc:    "正常系: stringのみ",
			program: `"hello"`,
			want: []token.Token{
				{
					Kind:     token.KindStr,
					ValueStr: "hello",
				},
			},
		},
		{
			desc:    "正常系: nilのみ",
			program: `nil`,
			want: []token.Token{
				{
					Kind:     token.KindNil,
					ValueNil: true,
				},
			},
		},
		{
			desc:    "正常系: symbolのみ",
			program: "exists?",
			want: []token.Token{
				{
					Kind:        token.KindSymbol,
					ValueSymbol: "exists?",
				},
			},
		},
		{
			desc:    "正常系: symbolのみ",
			program: "+",
			want: []token.Token{
				{
					Kind:        token.KindSymbol,
					ValueSymbol: "+",
				},
			},
		},
		{
			desc:    "正常系: 複数のatom",
			program: `+ 1 true "hello"`,
			want: []token.Token{
				{
					Kind:        token.KindSymbol,
					ValueSymbol: "+",
				},
				{
					Kind:     token.KindInt,
					ValueInt: 1,
				},
				{
					Kind:      token.KindBool,
					ValueBool: true,
				},
				{
					Kind:     token.KindStr,
					ValueStr: "hello",
				},
			},
		},
		{
			desc:    "正常系: 複数のatom (改行)",
			program: "+\n1\ntrue\n\"hello\"",
			want: []token.Token{
				{
					Kind:        token.KindSymbol,
					ValueSymbol: "+",
				},
				{
					Kind:     token.KindInt,
					ValueInt: 1,
				},
				{
					Kind:      token.KindBool,
					ValueBool: true,
				},
				{
					Kind:     token.KindStr,
					ValueStr: "hello",
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
						{
							Kind:        token.KindSymbol,
							ValueSymbol: "sym",
						},
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
						{
							Kind:        token.KindSymbol,
							ValueSymbol: "hello",
						},
						{
							Kind:     token.KindInt,
							ValueInt: 1,
						},
						{
							Kind:     token.KindStr,
							ValueStr: "world",
						},
						{
							Kind:     token.KindNil,
							ValueNil: true,
						},
						{
							Kind:      token.KindBool,
							ValueBool: true,
						},
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
						{
							Kind:        token.KindSymbol,
							ValueSymbol: "hello",
						},
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								{
									Kind:        token.KindSymbol,
									ValueSymbol: "foo",
								},
								{
									Kind:     token.KindInt,
									ValueInt: 1,
								},
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
						{
							Kind:        token.KindSymbol,
							ValueSymbol: "=",
						},
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								{
									Kind:        token.KindSymbol,
									ValueSymbol: "+",
								},
								{
									Kind:     token.KindInt,
									ValueInt: 1,
								},
								{
									Kind:     token.KindInt,
									ValueInt: 2,
								},
							},
						},
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								{
									Kind:        token.KindSymbol,
									ValueSymbol: "-",
								},
								{
									Kind:     token.KindInt,
									ValueInt: 3,
								},
								{
									Kind:     token.KindInt,
									ValueInt: 2,
								},
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
						{
							Kind:        token.KindSymbol,
							ValueSymbol: "+",
						},
						{
							Kind:     token.KindInt,
							ValueInt: 1,
						},
						{
							Kind: token.KindList,
							ValueList: []token.Token{
								{
									Kind:        token.KindSymbol,
									ValueSymbol: "-",
								},
								{
									Kind: token.KindList,
									ValueList: []token.Token{
										{
											Kind:        token.KindSymbol,
											ValueSymbol: "+",
										},
										{
											Kind:     token.KindInt,
											ValueInt: 1,
										},
										{
											Kind:     token.KindInt,
											ValueInt: 1,
										},
									},
								},
								{
									Kind: token.KindList,
									ValueList: []token.Token{
										{
											Kind:        token.KindSymbol,
											ValueSymbol: "*",
										},
										{
											Kind:     token.KindInt,
											ValueInt: 3,
										},
										{
											Kind: token.KindList,
											ValueList: []token.Token{
												{
													Kind:        token.KindSymbol,
													ValueSymbol: "/",
												},
												{
													Kind:     token.KindInt,
													ValueInt: 4,
												},
												{
													Kind:     token.KindInt,
													ValueInt: 5,
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
