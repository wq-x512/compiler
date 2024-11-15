package lexer

import (
	"regexp"
	"strings"
)

type Path struct {
	state     int
	ch        string
	nextState rune
	match     bool
	end       bool
}
type Token struct {
	Type  string
	Value string
}

// lexer Lexical analyzer
type LexerStruct struct {
	// Config
	CFG struct {
		ignoreTokens     []string
		ignoreValues     []string
		compressLineFeed bool // compress \n\n\n... to \n
	}
	// Input Stream Reader
	ISR struct {
		props struct {
			stream string // Character stream
			length int    // Length of character stream
			seq    int    // The sequence number of the character stream
		}
		propsChange struct {
			incrSeq   func()
			toDefault func()
		}
		before     func(string)
		after      func()
		nextChar   func() (rune, bool)
		isLastChar func() bool
		read       func()
	}
	DFA struct {
		result struct {
			matchs []string
			tokens []Token
		}
		resultChange struct {
			toDefault    func()
			pushToTokens func(Token)
			pushToMatchs func(string)
			filterTokens func()
			produceToken func()
		}
		state  int
		events struct {
			flowtoNextState  func(string, int)
			flowtoResetState func()
		}
	}
	resetDefault func()
	setConfig    func(interface{})
	Start        func(string)
}

var lexer LexerStruct

func init() {
	lexer = LexerStruct{
		CFG: struct {
			ignoreTokens     []string
			ignoreValues     []string
			compressLineFeed bool
		}{ignoreTokens: []string{}, ignoreValues: []string{WHITESPACE}, compressLineFeed: true},
		ISR: struct {
			props struct {
				stream string
				length int
				seq    int
			}
			propsChange struct {
				incrSeq   func()
				toDefault func()
			}
			before     func(string)
			after      func()
			nextChar   func() (rune, bool)
			isLastChar func() bool
			read       func()
		}{props: struct {
			stream string
			length int
			seq    int
		}{stream: "", length: 0, seq: 0}, propsChange: struct {
			incrSeq   func()
			toDefault func()
		}{incrSeq: func() {
			lexer.ISR.props.seq++
		}, toDefault: func() {
			lexer.ISR.props.stream = ""
			lexer.ISR.props.length = 0
			lexer.ISR.props.seq = 0
		}},
			before: func(stream string) {
				lexer.ISR.props.stream = stream
				if lexer.CFG.compressLineFeed {
					re := regexp.MustCompile(`\n+`)
					result := re.ReplaceAllString(lexer.ISR.props.stream, "\n")
					lexer.ISR.props.stream = strings.TrimSpace(result)
				}
				lexer.ISR.props.length = len(stream)
				lexer.ISR.props.seq = 0
			}, after: func() {
				lexer.DFA.resultChange.filterTokens()
			}, nextChar: func() (rune, bool) {
				seq := lexer.ISR.props.seq
				if seq <= lexer.ISR.props.length-1 {
					return rune(lexer.ISR.props.stream[seq]), true
				}
				return ' ', false
			}, isLastChar: func() bool {
				return lexer.ISR.props.seq == lexer.ISR.props.length-1
			}, read: func() {
				// TODO
				for ch, ok := lexer.ISR.nextChar(); ok; ch, ok = lexer.ISR.nextChar() {
					match := false
					end := false
					state := lexer.DFA.state
					nextState := FlowModel.getNextState(string(ch), state, lexer.DFA.result.matchs)
					if nextState != S_RESET {
						match = true
						if lexer.ISR.isLastChar() {
							end = true
						}
					}
					path := Path{
						state:     state,
						ch:        string(ch),
						nextState: nextState,
						match:     match,
						end:       end,
					}
					FlowModel.ResultChange.pathGrow(path)
					if match {
						lexer.ISR.propsChange.incrSeq()
						lexer.DFA.events.flowtoNextState(string(ch), int(nextState))
						if end {
							lexer.DFA.resultChange.produceToken()
						}
					} else {
						lexer.DFA.resultChange.produceToken()
						lexer.DFA.events.flowtoResetState()
					}
				}
			}},
		DFA: struct {
			result struct {
				matchs []string
				tokens []Token
			}
			resultChange struct {
				toDefault    func()
				pushToTokens func(Token)
				pushToMatchs func(string)
				filterTokens func()
				produceToken func()
			}
			state  int
			events struct {
				flowtoNextState  func(string, int)
				flowtoResetState func()
			}
		}{result: struct {
			matchs []string
			tokens []Token
		}{matchs: []string{}, tokens: []Token{}}, resultChange: struct {
			toDefault    func()
			pushToTokens func(Token)
			pushToMatchs func(string)
			filterTokens func()
			produceToken func()
		}{toDefault: func() {
			lexer.DFA.state = S_RESET
			lexer.DFA.result.tokens = nil
			lexer.DFA.result.matchs = nil
		}, pushToTokens: func(token Token) {
			lexer.DFA.result.tokens = append(lexer.DFA.result.tokens, token)
			lexer.DFA.result.matchs = nil
		}, pushToMatchs: func(ch string) {
			lexer.DFA.result.matchs = append(lexer.DFA.result.matchs, ch)
		}, filterTokens: func() {
			var tokens []Token
			for _, token := range lexer.DFA.result.tokens {
				if (!containsMap(lexer.CFG.ignoreValues, token.Value)) && (!containsMap(lexer.CFG.ignoreTokens, token.Type)) {
					tokens = append(tokens, token)
				}
			}
			lexer.DFA.result.tokens = tokens
		}, produceToken: func() {
			if len(lexer.DFA.result.matchs) != 0 {
				Value := strings.Join(lexer.DFA.result.matchs, "")
				Type := JudgeTokenType(rune(lexer.DFA.state), Value)
				token := Token{Type, Value}
				lexer.DFA.resultChange.pushToTokens(token)
			}
		}}, state: S_RESET, events: struct {
			flowtoNextState  func(string, int)
			flowtoResetState func()
		}{flowtoNextState: func(ch string, state int) {
			lexer.DFA.resultChange.pushToMatchs(ch)
			lexer.DFA.state = state
		}, flowtoResetState: func() {
			lexer.DFA.state = S_RESET
		}}},
		resetDefault: func() {
			FlowModel.ResultChange.toDefault()
			lexer.ISR.propsChange.toDefault()
			lexer.DFA.resultChange.toDefault()
		},
		setConfig: func(config interface{}) {
			// TODO
		},
		Start: func(s string) {
			lexer.resetDefault()
			lexer.ISR.before(s)
			lexer.ISR.read()
			lexer.ISR.after()
		},
	}
}
func GetToken() []Token {
	return lexer.DFA.result.tokens
}
func Start(s string) {
	lexer.Start(s)
}
