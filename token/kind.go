package token

type TokenKind int

const (
	KindUndefined TokenKind = iota
	KindBool
	KindInt
	KindStr
	KindNil
	KindSymbol
	KindList
)

func (t TokenKind) String() string {
	switch t {
	case KindBool:
		return "KindBool"
	case KindInt:
		return "KindInt"
	case KindStr:
		return "KindStr"
	case KindNil:
		return "KindNil"
	case KindSymbol:
		return "KindSymbol"
	case KindList:
		return "KindList"
	}
	return "KindUndefined"
}
