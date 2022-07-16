package parser

type Lexer struct {
	input string

	// the position of the current byte in the input
	currPos int

	// The position of the next byte to be read
	nextPos int

	// The current character being read
	char byte

	// the last token detected
	lastToken Token
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:     input,
		lastToken: Token{Type: EndOfFile, Literal: ""},
	}
	l.readChar()
	return l
}

// readChar reads the next input character and advances internal pointers
// currPos and nextPos
func (l *Lexer) readChar() {
	// check for end of input
	if l.nextPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPos]
	}
	l.currPos = l.nextPos
	l.nextPos += 1
}

// peekChar looks ahead to see what the next input character is
func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.char {
	case 0:
		// EOF
		tok.Type = EndOfFile
		tok.Literal = ""
	case '{':
		// left brace
		tok = newToken(LeftBrace, l.char)
	case '}':
		tok = newToken(RightBrace, l.char)
	case '[':
		tok = newToken(LeftSquareBrace, l.char)
	case ']':
		tok = newToken(RightSquareBrace, l.char)
	case ':':
		tok = newToken(Colon, l.char)
	case ',':
		tok = newToken(Comma, l.char)
	case '"':
		// look for closing quotes
		tok.Type = String
		tok.Literal = l.readString()
	}

	// read the next char
	l.readChar()
	l.lastToken = tok
	return tok
}

func (l *Lexer) readString() string {
	var strVal string = string(l.char)
	for l.peekChar() != '"' {
		l.readChar()
		strVal += string(l.char)
	}
	l.readChar() // read the closing quotes
	strVal += string(l.char)
	return strVal
}

func newToken(t TokenType, ch byte) Token {
	tok := Token{}
	tok.Type = t
	tok.Literal = string(ch)
	return tok
}
