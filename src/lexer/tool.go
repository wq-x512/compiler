package lexer

func containsMap(arr []string, elem string) bool {
	elementMap := make(map[string]bool)
	for _, v := range arr {
		elementMap[v] = true
	}
	return elementMap[elem]
}
func JudgeTokenTypeByValue(value string) string {
	if containsMap(Charset.Char.OPERATOR, value) {
		return "Operator"
	} else if containsMap(Charset.Char.SYMBOL, value) {
		return "Symbol"
	} else if containsMap(Charset.DoubleChar, value) {
		return "DoubleOperator"
	} else if containsMap(Charset.Keywords, value) {
		return "Keyword"
	}
	return ""
}
func JudgeTokenType(state rune, value string) string {
	tokentype := JudgeTokenTypeByValue(value)
	if len(tokentype) > 0 {
		return tokentype
	}
	if state == S_WHITESPACE {
		return "Whitespace"
	} else if state == S_LINEFEED {
		return "LineFeed"
	} else if state == S_IDENTIFIER {
		return "Identifier"
	} else if state == S_NUMBER {
		return "Number"
	} else if state == S_FLOAT {
		return "Float"
	} else if state == S_STRING_END {
		return "String"
	} else if state == S_CHAR_END {
		return "Char"
	}
	return "Unknown"
}

func GetFirstCharState(ch string) rune {
	if containsMap(Charset.DoubleCharFirstBitAnd, ch) {
		return S_DOUBLE_CHAR_FIRST_BIT_AND
	} else if containsMap(Charset.DoubleCharFirstBitOr, ch) {
		return S_DOUBLE_CHAR_FIRST_BIT_OR
	} else if containsMap(Charset.DoubleCharFirstGreat, ch) {
		return S_DOUBLE_CHAR_FIRST_GREAT
	} else if containsMap(Charset.DoubleCharFirstLess, ch) {
		return S_DOUBLE_CHAR_FIRST_LESS
	} else if containsMap(Charset.DoubleCharFirstNot, ch) {
		return S_DOUBLE_CHAR_FIRST_NOT
	} else if containsMap(Charset.DoubleCharFirstAssign, ch) {
		return S_DOUBLE_CHAR_FIRST_ASSIGN
	}
	return S_RESET
}
func GetSecondCharState(ch string) rune {
	if containsMap(Charset.DoubleCharSecondBitAnd, ch) {
		return S_DOUBLE_CHAR_SECOND_BIT_AND
	} else if containsMap(Charset.DoubleCharSecondBitOr, ch) {
		return S_DOUBLE_CHAR_SECOND_BIT_OR
	} else if containsMap(Charset.DoubleCharSecondGreat, ch) {
		return S_DOUBLE_CHAR_SECOND_GREAT
	} else if containsMap(Charset.DoubleCharSecondLess, ch) {
		return S_DOUBLE_CHAR_SECOND_LESS
	} else if containsMap(Charset.DoubleCharSecondAssign, ch) {
		return S_DOUBLE_CHAR_SECOND_ASSIGN
	}
	return S_RESET
}
func IsOperatorChar(char string) bool {
	return containsMap(Charset.Char.OPERATOR, char)
}
func IsSymbolChar(char string) bool {
	return containsMap(Charset.Char.SYMBOL, char)
}
func IsWhitespaceChar(char string) bool {
	return containsMap(Charset.Char.WHITESPACE, char)
}
func IsLineFeedChar(char string) bool {
	return containsMap(Charset.Char.LINEFEED, char)
}
func IsFirstCharOfDoubleChar(char string) bool {
	return GetFirstCharState(char) > 0
}
func IsSecondCharOfDoubleChar(char string) bool {
	return GetSecondCharState(char) > 0
}
func IsAlphabetChar(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}
func IsNumberChar(char rune) bool {
	return char >= '0' && char <= '9'
}
func IsChinese(char rune) bool {
	return char >= '\u4e00' && char <= '\u9fff'
}
func IsInStates(state int, states []int) bool {
	for _, v := range states {
		if state == v {
			return true
		}
	}
	return false
}
