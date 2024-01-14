package interpreter

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tejasdeepakmasne/bfit/lexer"
)

type Machine struct {
	Tape    []uint8
	Program []lexer.Token
	pc      int
	mp      int
}

func Interpret(program []lexer.Token) {
	initialTape := make([]uint8, 65536)
	mc := Machine{Tape: initialTape, Program: program, pc: 0, mp: 0}
	it := 0
	for mc.pc < len(mc.Program) {
		it++
		//fmt.Printf("it: %d, pc: %d, mp: %d, val: %v\n", it, mc.pc, mc.mp, (mc.Tape[mc.mp]))
		switch mc.Program[mc.pc].Type {
		case lexer.MOVE_LEFT:
			if mc.mp >= len(mc.Tape)-1 {
				additionalTape := make([]uint8, 2)
				mc.Tape = append(mc.Tape, additionalTape...)
				mc.mp += 1
				mc.pc += 1
			} else {
				mc.mp += 1
				mc.pc += 1
			}
		case lexer.MOVE_RIGHT:
			if mc.mp <= 0 {
				additionalTape := make([]uint8, 2)
				mc.Tape = append(additionalTape, mc.Tape...)
				mc.mp -= 1
				mc.pc += 1
			} else {
				mc.mp -= 1
				mc.pc += 1
			}
		case lexer.INCREMENT:
			if mc.Tape[mc.mp] == 255 {
				mc.Tape[mc.mp] = 0
				mc.pc += 1
			} else {
				mc.Tape[mc.mp] += 1
				mc.pc += 1
			}
		case lexer.DECREMENT:
			if mc.Tape[mc.mp] == 0 {
				mc.Tape[mc.mp] = 255
				mc.pc += 1
			} else {
				mc.Tape[mc.mp] -= 1
				mc.pc += 1
			}
		case lexer.WRITE:
			fmt.Printf("%s", string(mc.Tape[mc.mp]%128))
			mc.pc += 1
		case lexer.READ:
			reader := bufio.NewReader(os.Stdin)
			var err error
			mc.Tape[mc.mp], err = reader.ReadByte()
			if err != nil {
				panic(err)
			}
			mc.pc += 1
		case lexer.JMP_IF_ZERO:
			if mc.Tape[mc.mp] != 0 {
				mc.pc += 1
				//fmt.Printf("JMP false value %v\n", string(mc.Tape[mc.mp]))
			} else {
				//fmt.Printf("JMP pc initial: %d\n", mc.pc)

				//fmt.Printf("JMP true value %v\n", string(mc.Tape[mc.mp]))
				mc.pc = mc.Program[mc.pc].PositionChange
				//fmt.Printf("JMP pc change: %d\n", mc.pc)
			}
		case lexer.JMP_UNLESS_ZERO:
			if mc.Tape[mc.mp] != 0 {
				//fmt.Printf("JMP_UNLESS pc initial: %d\n", mc.pc)
				//fmt.Printf("JMP_UNLESS true value %v\n", string(mc.Tape[mc.mp]))
				mc.pc = mc.Program[mc.pc].PositionChange
				//fmt.Printf("JMP_UNLESS pc change: %d\n", mc.pc)
			} else {
				mc.pc += 1
				//fmt.Printf("JMP_UNLESS false value %v\n", string(mc.Tape[mc.mp]))
			}
		}
	}
}
