package lexer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	input := []byte{'>', '<', '+', '-', ',', '.', '[', ']'}
	tests := []struct {
		expectedType           TokenType
		expectedPositionChange int
	}{
		{MOVE_LEFT, 1},
		{MOVE_RIGHT, -1},
		{INCREMENT, 0},
		{DECREMENT, 0},
		{READ, 0},
		{WRITE, 0},
		{JMP_IF_ZERO, 7},
		{JMP_UNLESS_ZERO, 6},
	}

	tokensReceived := GenerateTokens(input)
	//fmt.Println(tokensReceived)
	for i, tt := range tests {
		if tokensReceived[i].Type != tt.expectedType {
			t.Fatalf("tests[%d] wrong TokenType. Expected : %q Got : %q", i, tt.expectedType, tokensReceived[i].Type)
		} else {
			fmt.Printf("type : %v", tokensReceived[i].Type)
		}
		if tokensReceived[i].PositionChange != tt.expectedPositionChange {
			t.Fatalf("tests[%d] wrong PositionChange. Expected : %d Got : %d", i, tt.expectedPositionChange, tokensReceived[i].PositionChange)

		} else {
			fmt.Printf("positionChange : %v\n", tokensReceived[i].PositionChange)
		}
	}
}
