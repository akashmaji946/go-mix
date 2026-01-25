package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// represents a test case for ConsumeTokens
// Input: source code
// ExpectedTokens: list of expected tokens
type TestConsumeToken struct {
	Input          string
	ExpectedTokens []Token
}

// TestNewLexer_ConsumeTokens tests the ConsumeTokens method of the Lexer
func TestNewLexer_ConsumeTokens(t *testing.T) {

	tests := []TestConsumeToken{
		{
			Input: ` 123 + 2   31 - 12 `,
			ExpectedTokens: []Token{
				NewToken(NUMBER_ID, "123"),
				NewToken(PLUS_OP, "+"),
				NewToken(NUMBER_ID, "2"),
				NewToken(NUMBER_ID, "31"),
				NewToken(MINUS_OP, "-"),
				NewToken(NUMBER_ID, "12"),
			},
		},
		{
			Input: ` { } + []  abc - a12 `,
			ExpectedTokens: []Token{
				NewToken(LEFT_BRACE, "{"),
				NewToken(RIGHT_BRACE, "}"),
				NewToken(PLUS_OP, "+"),
				NewToken(LEFT_BRACKET, "["),
				NewToken(RIGHT_BRACKET, "]"),
				NewToken(IDENTIFIER_ID, "abc"),
				NewToken(MINUS_OP, "-"),
				NewToken(IDENTIFIER_ID, "a12"),
			},
		},
		{
			Input: ` <=  + 2   {31} - 12 __a19bcd_aa90`,
			ExpectedTokens: []Token{
				NewToken(LE_OP, "<="),
				NewToken(PLUS_OP, "+"),
				NewToken(NUMBER_ID, "2"),
				NewToken(LEFT_BRACE, "{"),
				NewToken(NUMBER_ID, "31"),
				NewToken(RIGHT_BRACE, "}"),
				NewToken(MINUS_OP, "-"),
				NewToken(NUMBER_ID, "12"),
				NewToken(IDENTIFIER_ID, "__a19bcd_aa90"),
			},
		},
		{
			Input: `"This is a long string  " nowAnIdentifier_234 "12"`,
			ExpectedTokens: []Token{
				NewToken(STRING_LIT, "This is a long string  "),
				NewToken(IDENTIFIER_ID, "nowAnIdentifier_234"),
				NewToken(STRING_LIT, "12"),
			},
		},

		{
			Input: `func new if else then for abc123 "hello!" __KEY__`,
			ExpectedTokens: []Token{
				NewToken(FUNC_KEY, "func"),
				NewToken(NEW_KEY, "new"),
				NewToken(IF_KEY, "if"),
				NewToken(ELSE_KEY, "else"),
				NewToken(IDENTIFIER_ID, "then"),
				NewToken(FOR_KEY, "for"),
				NewToken(IDENTIFIER_ID, "abc123"),
				NewToken(STRING_LIT, "hello!"),
				NewToken(IDENTIFIER_ID, "__KEY__"),
			},
		},

		{
			Input: `
			func main(args, argv) {
				var a = args[0];
				var b = argv[0];
				if (a <= 0){
					return a + b;
				} else{
					var f = 1;
					while (f < b){
						f = f * a + 2;
					}
					return f;
				}
			}
			`,
			ExpectedTokens: []Token{
				NewToken(FUNC_KEY, "func"),
				NewToken(IDENTIFIER_ID, "main"),
				NewToken(LEFT_PAREN, "("),
				NewToken(IDENTIFIER_ID, "args"),
				NewToken(COMMA_DELIM, ","),
				NewToken(IDENTIFIER_ID, "argv"),
				NewToken(RIGHT_PAREN, ")"),
				NewToken(LEFT_BRACE, "{"),
				NewToken(VAR_KEY, "var"),
				NewToken(IDENTIFIER_ID, "a"),
				NewToken(ASSIGN_OP, "="),
				NewToken(IDENTIFIER_ID, "args"),
				NewToken(LEFT_BRACKET, "["),
				NewToken(NUMBER_ID, "0"),
				NewToken(RIGHT_BRACKET, "]"),
				NewToken(SEMICOLON_DELIM, ";"),
				NewToken(VAR_KEY, "var"),
				NewToken(IDENTIFIER_ID, "b"),
				NewToken(ASSIGN_OP, "="),
				NewToken(IDENTIFIER_ID, "argv"),
				NewToken(LEFT_BRACKET, "["),
				NewToken(NUMBER_ID, "0"),
				NewToken(RIGHT_BRACKET, "]"),
				NewToken(SEMICOLON_DELIM, ";"),
				NewToken(IF_KEY, "if"),
				NewToken(LEFT_PAREN, "("),
				NewToken(IDENTIFIER_ID, "a"),
				NewToken(LE_OP, "<="),
				NewToken(NUMBER_ID, "0"),
				NewToken(RIGHT_PAREN, ")"),
				NewToken(LEFT_BRACE, "{"),
				NewToken(RETURN_KEY, "return"),
				NewToken(IDENTIFIER_ID, "a"),
				NewToken(PLUS_OP, "+"),
				NewToken(IDENTIFIER_ID, "b"),
				NewToken(SEMICOLON_DELIM, ";"),
				NewToken(RIGHT_BRACE, "}"),
				NewToken(ELSE_KEY, "else"),
				NewToken(LEFT_BRACE, "{"),
				NewToken(VAR_KEY, "var"),
				NewToken(IDENTIFIER_ID, "f"),
				NewToken(ASSIGN_OP, "="),
				NewToken(NUMBER_ID, "1"),
				NewToken(SEMICOLON_DELIM, ";"),
				NewToken(WHILE_KEY, "while"),
				NewToken(LEFT_PAREN, "("),
				NewToken(IDENTIFIER_ID, "f"),
				NewToken(LT_OP, "<"),
				NewToken(IDENTIFIER_ID, "b"),
				NewToken(RIGHT_PAREN, ")"),
				NewToken(LEFT_BRACE, "{"),
				NewToken(IDENTIFIER_ID, "f"),
				NewToken(ASSIGN_OP, "="),
				NewToken(IDENTIFIER_ID, "f"),
				NewToken(MUL_OP, "*"),
				NewToken(IDENTIFIER_ID, "a"),
				NewToken(PLUS_OP, "+"),
				NewToken(NUMBER_ID, "2"),
				NewToken(SEMICOLON_DELIM, ";"),
				NewToken(RIGHT_BRACE, "}"),
				NewToken(RETURN_KEY, "return"),
				NewToken(IDENTIFIER_ID, "f"),
				NewToken(SEMICOLON_DELIM, ";"),
				NewToken(RIGHT_BRACE, "}"),
				NewToken(RIGHT_BRACE, "}"),
			},
		},
	}

	for _, test := range tests {
		lex := NewLexer(test.Input)

		gotTokens := lex.ConsumeTokens()

		// must: length match
		assert.Equal(t, len(test.ExpectedTokens), len(gotTokens))
		// must: token to token match
		for i, token := range test.ExpectedTokens {
			assert.Equal(t, token.Type, gotTokens[i].Type)
			assert.Equal(t, token.Literal, gotTokens[i].Literal)
		}
	}

}
