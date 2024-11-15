package lexer

const (
	// DFA State
	S_RESET = iota
	S_OPERATOR
	S_SYMBOL
	S_DOUBLE_CHAR_FIRST_BIT_AND
	S_DOUBLE_CHAR_SECOND_BIT_AND
	S_DOUBLE_CHAR_FIRST_BIT_OR
	S_DOUBLE_CHAR_SECOND_BIT_OR
	S_DOUBLE_CHAR_FIRST_GREAT
	S_DOUBLE_CHAR_SECOND_GREAT
	S_DOUBLE_CHAR_FIRST_LESS
	S_DOUBLE_CHAR_SECOND_LESS
	S_DOUBLE_CHAR_FIRST_NOT
	S_DOUBLE_CHAR_FIRST_ASSIGN
	S_DOUBLE_CHAR_SECOND_ASSIGN
	S_WHITESPACE
	S_LINEFEED
	S_IDENTIFIER
	S_NUMBER
	S_FLOAT
	S_STRING
	S_STRING_ESCAPE
	S_STRING_END
	S_CHAR
	S_CHAR_END
	S_END
)

type FlowModelStruct struct {
	Result struct {
		Path []interface{}
	}
	ResultChange struct {
		pathGrow  func(path interface{})
		toDefault func()
	}
	getNextState func(ch string, state int, matchs []string) rune
}

var FlowModel FlowModelStruct

func init() {
	FlowModel = FlowModelStruct{
		Result: struct{ Path []interface{} }{Path: nil},
		ResultChange: struct {
			pathGrow  func(path interface{})
			toDefault func()
		}{pathGrow: func(path interface{}) {
			FlowModel.Result.Path = append(FlowModel.Result.Path, path)
		}, toDefault: func() {
			FlowModel.Result.Path = []interface{}{}
		}},
		getNextState: func(ch string, state int, matchs []string) rune {
			if IsInStates(state, []int{S_STRING, S_STRING_ESCAPE}) {
				if state == S_STRING && ch == ESCAPE {
					return S_STRING_ESCAPE
				}
				if state == S_STRING_ESCAPE {
					return S_STRING
				}
				if ch != DOUBLE_QUOTATION {
					return S_STRING
				}
				return S_STRING_END
			}
			if IsInStates(state, []int{S_CHAR}) {
				if len(matchs) == 1 {
					return S_CHAR
				}
				if len(matchs) == 2 && ch == QUOTATION {
					return S_CHAR_END
				}
				return S_RESET
			}
			if IsAlphabetChar([]rune(ch)[0]) {
				if IsInStates(state, []int{S_RESET, S_IDENTIFIER}) {
					return S_IDENTIFIER
				}
			} else if IsNumberChar([]rune(ch)[0]) {
				if IsInStates(state, []int{S_RESET, S_NUMBER}) {
					return S_NUMBER
				}
				if IsInStates(state, []int{S_FLOAT}) {
					return S_FLOAT
				}
				if IsInStates(state, []int{S_IDENTIFIER}) {
					return S_IDENTIFIER
				}
			} else if IsOperatorChar(ch) {
				if IsFirstCharOfDoubleChar(ch) {
					if IsInStates(state, []int{S_RESET}) {
						return GetFirstCharState(ch)
					}
				}
				if IsSecondCharOfDoubleChar(ch) {
					if state == S_DOUBLE_CHAR_FIRST_BIT_AND && ch == BIT_AND {
						return S_DOUBLE_CHAR_SECOND_BIT_AND
					}
					if state == S_DOUBLE_CHAR_FIRST_BIT_OR && ch == BIT_OR {
						return S_DOUBLE_CHAR_SECOND_BIT_OR
					}
					if state == S_DOUBLE_CHAR_FIRST_GREAT && ch == GREAT {
						return S_DOUBLE_CHAR_SECOND_GREAT
					}
					if state == S_DOUBLE_CHAR_FIRST_LESS && ch == LESS {
						return S_DOUBLE_CHAR_SECOND_LESS
					}
					if state == S_DOUBLE_CHAR_FIRST_GREAT && ch == ASSIGN {
						return S_DOUBLE_CHAR_SECOND_ASSIGN
					}
					if state == S_DOUBLE_CHAR_FIRST_LESS && ch == ASSIGN {
						return S_DOUBLE_CHAR_SECOND_ASSIGN
					}
					if state == S_DOUBLE_CHAR_FIRST_NOT && ch == ASSIGN {
						return S_DOUBLE_CHAR_SECOND_ASSIGN
					}
					if state == S_DOUBLE_CHAR_FIRST_ASSIGN && ch == ASSIGN {
						return S_DOUBLE_CHAR_SECOND_ASSIGN
					}
				}
				if IsInStates(state, []int{S_RESET}) {
					return S_OPERATOR
				}
			} else if IsSymbolChar(ch) {
				if IsInStates(state, []int{S_RESET}) {
					if ch == QUOTATION {
						return S_CHAR
					}
					if ch == DOUBLE_QUOTATION {
						return S_STRING
					}
					return S_SYMBOL
				}
				if IsInStates(state, []int{S_NUMBER}) && ch == DOT {
					return S_FLOAT
				}
			} else if IsWhitespaceChar(ch) {
				if IsInStates(state, []int{S_RESET}) {
					return S_WHITESPACE
				}
			} else if IsLineFeedChar(ch) {
				if IsInStates(state, []int{S_RESET}) {
					return S_LINEFEED
				}
			} else {
				if IsInStates(state, []int{S_RESET, S_IDENTIFIER}) {
					return S_IDENTIFIER
				}
			}
			return S_RESET
		},
	}
}
