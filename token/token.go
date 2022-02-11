package token

type Token struct {
	Kind        TokenKind
	ValueBool   bool
	ValueInt    int64
	ValueStr    string
	ValueNil    bool
	ValueSymbol string
	ValueList   []Token
}

func NewBoolToken(b bool) Token {
	return Token{
		Kind:      KindBool,
		ValueBool: b,
	}
}

func NewIntToken(i int64) Token {
	return Token{
		Kind:     KindInt,
		ValueInt: i,
	}
}

func NewStrToken(s string) Token {
	return Token{
		Kind:     KindStr,
		ValueStr: s,
	}
}

func NewNilToken() Token {
	return Token{
		Kind:     KindNil,
		ValueNil: true,
	}
}

func NewSymbolToken(s string) Token {
	return Token{
		Kind:        KindSymbol,
		ValueSymbol: s,
	}
}

func NewListToken() Token {
	return Token{
		Kind: KindList,
	}
}
