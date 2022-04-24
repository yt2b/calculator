package token

type Type string

type Token struct {
	Type  Type
	Value string
}

// トークンの種類
const (
	NUMBER   Type = "NUMBER"
	PLUS          = "+"
	MINUS         = "-"
	ASTERISK      = "*"
	SLASH         = "/"
	LPAREN        = "("
	RPAREN        = ")"
	EOF           = "\x00"
	UNKNOWN       = "UNKNOWN"
)
