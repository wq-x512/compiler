package lexer

type Tag uint32

const (
	// Simple Operator
	PLUS     = "+"
	SUBTRACT = "-"
	TIMES    = "*"
	DIVIDE   = "/"
	MODE     = "%"

	// Logical Operator
	NOT = "!"

	// Compare Operator
	GREAT  = ">"
	LESS   = "<"
	ASSIGN = "="

	GREAT_EQUAL = ">="
	LESS_EQUAL  = "<="
	EQUAL       = "=="
	NOT_EQUAL   = "!="

	// Bit Operator
	BIT_AND = "&"
	BIT_OR  = "|"
	BIT_NOT = "~"
	BIT_XOR = "^"

	CONDITION_AND   = "&&"
	CONDITION_OR    = "||"
	BIT_LEFT_SHIFT  = "<<"
	BIT_RIGHT_SHIFT = ">>"

	// Whitespace mark
	WHITESPACE = " "
	LINEFEED   = "\n"

	// Bracket mark
	LEFT_BRACKET         = "("
	RIGHT_BRACKET        = ")"
	LEFT_BRACE           = "{"
	RIGHT_BRACE          = "}"
	LEFT_SQUARE_BRACKET  = "["
	RIGHT_SQUARE_BRACKET = "]"

	// Other mark
	DOT              = "."
	COMMA            = ","
	SEMICOLON        = ";"
	QUOTATION        = "'"
	DOUBLE_QUOTATION = "\""
	SHARP            = "#"
	QUESTION         = "?"
	COLON            = ":"
	ESCAPE           = "\\"
	// TOKEN
	T_OPERATOR        = 1
	T_DOUBLE_OPERATOR = 2
	T_SYMBOL          = 3
	T_WHITESPACE      = 4
	T_LINEFEED        = 5
	KEYWORD           = 6
	T_IDENTIFIER      = 7
	T_NUMBER          = 8
	T_FLOAT           = 9
	T_STRING          = 10
	T_CHAR            = 11
	T_UNKNOWN         = 12
)

type CHARSET struct {
	Char                   CHAR
	DoubleChar             []string
	DoubleCharFirstBitAnd  []string
	DoubleCharSecondBitAnd []string
	DoubleCharFirstBitOr   []string
	DoubleCharSecondBitOr  []string
	DoubleCharFirstGreat   []string
	DoubleCharSecondGreat  []string
	DoubleCharFirstLess    []string
	DoubleCharSecondLess   []string
	DoubleCharFirstNot     []string
	DoubleCharFirstAssign  []string
	DoubleCharSecondAssign []string
	Keywords               []string
}
type CHAR struct {
	OPERATOR   []string
	SYMBOL     []string
	WHITESPACE []string
	LINEFEED   []string
}

var Charset CHARSET

func init() {
	Charset = CHARSET{
		Char: CHAR{
			OPERATOR:   []string{PLUS, SUBTRACT, TIMES, DIVIDE, MODE, NOT, GREAT, LESS, ASSIGN, BIT_AND, BIT_OR, BIT_NOT, BIT_XOR},
			SYMBOL:     []string{LEFT_BRACKET, RIGHT_BRACKET, LEFT_SQUARE_BRACKET, RIGHT_SQUARE_BRACKET, LEFT_BRACE, RIGHT_BRACE, DOT, COMMA, SEMICOLON, QUOTATION, DOUBLE_QUOTATION, SHARP, QUESTION, COLON},
			WHITESPACE: []string{WHITESPACE},
			LINEFEED:   []string{LINEFEED},
		},
		DoubleChar:             []string{GREAT_EQUAL, LESS_EQUAL, EQUAL, NOT_EQUAL, CONDITION_AND, CONDITION_OR, BIT_LEFT_SHIFT, BIT_RIGHT_SHIFT},
		DoubleCharFirstBitAnd:  []string{BIT_AND},
		DoubleCharSecondBitAnd: []string{BIT_AND},
		DoubleCharFirstBitOr:   []string{BIT_OR},
		DoubleCharSecondBitOr:  []string{BIT_OR},
		DoubleCharFirstGreat:   []string{GREAT},
		DoubleCharSecondGreat:  []string{GREAT},
		DoubleCharFirstLess:    []string{LESS},
		DoubleCharSecondLess:   []string{LESS},
		DoubleCharFirstNot:     []string{NOT},
		DoubleCharFirstAssign:  []string{ASSIGN},
		DoubleCharSecondAssign: []string{ASSIGN},
		Keywords: []string{
			"char", "int", "short", "long", "float", "double", "sizeof", "signed", "unsigned",
			"if", "else", "while", "for", "do", "break", "continue", "goto", "main",
			"void", "return", "switch", "case", "default",
			"const", "static", "auto", "extern", "register",
			"struct", "union", "enum", "typedef",
			"include",
		},
	}
}
