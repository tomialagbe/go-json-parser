package parser

import "testing"

func TestLexer(t *testing.T) {

	testSuites := map[string][]struct {
		Type    TokenType
		Literal string
	}{
		`{"a": "aVal"}`: {
			{LeftBrace, "{"},
			{String, `"a"`},
			{Colon, ":"},
			{String, `"aVal"`},
			{RightBrace, "}"},
			{EndOfFile, ""},
		},
		`{"keyOne": "keyOneVal", "keyTwo": "keyTwoVal"}`: {
			{LeftBrace, "{"},
			{String, `"keyOne"`},
			{Colon, ":"},
			{String, `"keyOneVal"`},
			{Comma, ","},
			{String, `"keyTwo"`},
			{Colon, ":"},
			{String, `"keyTwoVal"`},
			{RightBrace, "}"},
			{EndOfFile, ""},
		},
		`["item1", "item2"]`: {
			{LeftSquareBrace, "["},
			{String, `"item1"`},
			{Comma, ","},
			{String, `"item2"`},
			{RightSquareBrace, "]"},
		},
	}

	for input, tests := range testSuites {
		lexer := NewLexer(input)
		for idx, test := range tests {
			token := lexer.NextToken()
			if token.Type != test.Type {
				t.Fatalf("Test[%d] failed. Expected token type %v, got %v", idx, test.Type, token.Type)
			}

			if token.Literal != test.Literal {
				t.Fatalf("Test[%d] failed. Expected token literal %v, got %v", idx, test.Literal, token.Literal)
			}
		}
	}

}
