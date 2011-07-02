package main

import (
	"fmt"
	"strings"
	"strconv"
)

func RunProgramString(code string, logoperandstack bool, initmems string) {
	RunProgram([]byte(code), logoperandstack, initmems)
}

func RunProgram(code []byte, logoperandstack bool, initmems string) {
	eip, oldeip := 0, 0
	memory := NewRam(MEMORY_SIZE)
	operandstack := NewStack(OPERANDSTACK_SIZE)
	callstack := NewStack(CALLSTACK_SIZE)
	
	if initmems != "" {
		initmem := strings.Split(initmems, ",", -1)
		for i, s := range(initmem) {
			m, err := strconv.Atoi(s)
			if err == nil {		
				memory.Set(i, m)
			}
		}
	}
	
	var log []string
	if logoperandstack {
		log = make([]string, 0)
	}
	
	defer func() {
		println()
	
        if e := recover(); e != nil {
			fmt.Printf("   %s caused by %c with eip %d\n   operandstack=%d, callstack=%d\n\n", e.(string), code[oldeip], oldeip, operandstack.Sp, callstack.Sp)
        }
		
		if logoperandstack {
			for _, s := range(log) {
				println(s)
			}
		}
    }()
	
	for eip < len(code) {
		if eip < 0 || eip >= len(code) {
			panic("eip out of code range")
		}
	
		oldeip = eip
		
		bc := Bytecode(code[eip])
		
		eip++
		
		switch(bc) {
			case NOP:
			case NOP2:
			case PINT:
				print(operandstack.Pop())
			case PCHR:
				print(string(operandstack.Pop()))
			case P0:
				operandstack.Push(0)
			case P1:
				operandstack.Push(1)
			case P2:
				operandstack.Push(2)
			case P3:
				operandstack.Push(3)	
			case P4:
				operandstack.Push(4)	
			case P5:
				operandstack.Push(5)
			case P6:
				operandstack.Push(6)
			case P7:
				operandstack.Push(7)
			case P8:
				operandstack.Push(8)
			case P9:
				operandstack.Push(9)
			case ADD:
				S0, S1 := operandstack.Pop(), operandstack.Pop()
				operandstack.Push(S1 + S0)
			case MUL:
				S0, S1 := operandstack.Pop(), operandstack.Pop()
				operandstack.Push(S1 * S0)
			case SUB:
				S0, S1 := operandstack.Pop(), operandstack.Pop()
				operandstack.Push(S1 - S0)
			case DIV:
				S0, S1 := operandstack.Pop(), operandstack.Pop()
				operandstack.Push(S1 / S0)
			case CMP:
				S0, S1 := operandstack.Pop(), operandstack.Pop()
				switch {
					case S0 > S1:
						operandstack.Push(-1)
					case S0 < S1:
						operandstack.Push(1)
					default:
						operandstack.Push(0)
				}
			case JMP:
				S0 := operandstack.Pop()
				t := eip + S0
				if t < 0 || t >= len(code) {
					panic("jump out of code range")
				}
				eip += S0
			case JNE:
				S0, S1 := operandstack.Pop(), operandstack.Pop()
				if S1 == 0 {
					t := eip + S0
					if t < 0 || t > len(code) {
						panic("jump out of code range")
					}
					eip += S0
				}
			case CALL:
				callstack.Push(eip)
				eip = operandstack.Pop()
			case RET:
				eip = callstack.Pop()
			case PEEK:
				operandstack.Push(memory.Get(operandstack.Pop()))
			case POKE:
				memory.Set(operandstack.Pop(), operandstack.Pop())
			case PICK:
				operandstack.Push(operandstack.Get(operandstack.Sp - operandstack.Pop() - 1))
			case ROLL:
				S0 := operandstack.Pop()
				t := operandstack.Sp - S0 - 1
				v := operandstack.Get(t)
				operandstack.Del(t)
				operandstack.Push(v)
			case DROP:
				operandstack.Pop()
			case END:
				eip = len(code)
			default:
				panic("unknown bytecode")
		}
		
		if logoperandstack {
			log = append(log, fmt.Sprintf("%c - %v", code[oldeip], operandstack.ValidValues()))
		}
	}
}