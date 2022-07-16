package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LeftBrace        = "{"
	RightBrace       = "}"
	LeftSquareBrace  = "["
	RightSquareBrace = "]"
	Colon            = ":"
	Comma            = ","
	Int              = "Int"
	Float            = "Float"
	String           = "String"
	Bool             = "Bool"

	// Helpers
	EndOfFile = "EOF"
)
