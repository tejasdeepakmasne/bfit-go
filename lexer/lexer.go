package lexer

import (
	"errors"
)

type TokenType byte

const (
	MOVE_LEFT       = '>'
	MOVE_RIGHT      = '<'
	INCREMENT       = '+'
	DECREMENT       = '-'
	JMP_IF_ZERO     = '['
	JMP_UNLESS_ZERO = ']'
	READ            = ','
	WRITE           = '.'
)

type Token struct {
	Type           TokenType
	PositionChange int
}

type Scanner struct {
	Input        []byte
	Position     int
	NextPosition int
	Char         byte
	BracketStack []int
}

// I dont know where else to put this but this is a simple stack
func (sc *Scanner) Push(index int) {
	sc.BracketStack = append(sc.BracketStack, index)
}

func (sc *Scanner) Pop() (int, error) {
	stackLen := len(sc.BracketStack)
	var topElement int
	if stackLen > 0 {
		topElement = sc.BracketStack[stackLen-1]
		stackLen -= 1
		sc.BracketStack = sc.BracketStack[:stackLen]
		return topElement, nil
	} else {
		return -1, errors.New("stack Empty nothing to pop")
	}
}

func (sc *Scanner) Top() (int, error) {
	if len(sc.BracketStack) > 0 {
		return sc.BracketStack[len(sc.BracketStack)-1], nil
	} else {
		return -1, errors.New("stack empty no top element")
	}
}

func (sc *Scanner) readChar() bool {
	//fmt.Println(sc)
	if sc.NextPosition > len(sc.Input) {
		return false
	} else {
		sc.Char = sc.Input[sc.Position]
	}
	sc.Position = sc.NextPosition
	sc.NextPosition += 1
	return true
}

var leftBracketIndices map[int]int = make(map[int]int)
var rightBracketIndices map[int]int = make(map[int]int)

func (sc *Scanner) matchBrackets() {
	for i, ch := range sc.Input {
		if ch == '[' {
			sc.Push(i)
		}
		if ch == ']' {
			index, err := sc.Pop()
			//fmt.Printf("top element index : %d", index)
			if err != nil {
				panic(err)
			}
			leftBracketIndices[index] = i
			rightBracketIndices[i] = index

		}
	}
}

func GenerateTokens(source []byte) []Token {
	var generatedTokens []Token
	sc := Scanner{Input: source, NextPosition: 1, Position: 0}
	sc.matchBrackets()
	//fmt.Println(leftBracketIndices)
	//fmt.Println(rightBracketIndices)
	for sc.readChar() {
		switch sc.Char {
		case '>':
			generatedTokens = append(generatedTokens, Token{MOVE_LEFT, 1})
		case '<':
			generatedTokens = append(generatedTokens, Token{MOVE_RIGHT, -1})
		case '+':
			generatedTokens = append(generatedTokens, Token{INCREMENT, 0})
		case '-':
			generatedTokens = append(generatedTokens, Token{DECREMENT, 0})
		case ',':
			generatedTokens = append(generatedTokens, Token{READ, 0})
		case '.':
			generatedTokens = append(generatedTokens, Token{WRITE, 0})
		case '[':
			//fmt.Println(sc.Position)
			//fmt.Println(leftBracketIndices[sc.NextPosition-1])
			generatedTokens = append(generatedTokens, Token{JMP_IF_ZERO, leftBracketIndices[sc.Position-1]})
		case ']':
			generatedTokens = append(generatedTokens, Token{JMP_UNLESS_ZERO, rightBracketIndices[sc.Position-1]})
		default:
			sc.readChar()
		}
	}

	return generatedTokens
}
