package gohackvm

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	MEMORY_SIZE = 16384
	OPERANDSTACK_SIZE = 256
	CALLSTACK_SIZE = 256
)

type HackVM struct {
	eip int
	memory *Ram
	operandstack *Stack
	callstack *Stack
	logoperandstack bool
}

func NewHackVM(logoperandstack bool) *HackVM {
	return &HackVM{0, 
		NewRam(MEMORY_SIZE), 
		NewStack(OPERANDSTACK_SIZE), 
		NewStack(CALLSTACK_SIZE), 
		logoperandstack}
}

func (vm* HackVM) SetInitialMemory(initmems string) {
	if initmems != "" {
		initmem := strings.Split(initmems, ",")
		for i, s := range(initmem) {
			m, err := strconv.Atoi(s)
			if err == nil {		
				vm.memory.Set(i, m)
			}
		}
	}
}

func (vm* HackVM) RunProgramString(code string) {
	vm.RunProgram([]byte(code))
}

func (vm* HackVM) RunProgram(code []byte) {
	oldeip := 0

	var log []string
	if vm.logoperandstack {
		log = make([]string, 0)
	}
	
	defer func() {
		println()
	
        if e := recover(); e != nil {
			fmt.Printf("   %s caused by %c with eip %d\n   operandstack=%d, callstack=%d\n\n", 
				e.(string), 
				code[oldeip], 
				oldeip, 
				vm.operandstack.Sp, 
				vm.callstack.Sp)
        }
		
		if vm.logoperandstack {
			for _, s := range(log) {
				println(s)
			}
		}
    }()
	
	for vm.eip < len(code) {
		if vm.eip < 0 || vm.eip >= len(code) {
			panic("eip out of code range")
		}
	
		oldeip = vm.eip
		bc := Bytecode(code[vm.eip])
		vm.eip++
		
		switch(bc) {
			case NOP:
			case NOP2:
			case PINT:
				print(vm.operandstack.Pop())
			case PCHR:
				print(string(vm.operandstack.Pop()))
			case P0:
				vm.operandstack.Push(0)
			case P1:
				vm.operandstack.Push(1)
			case P2:
				vm.operandstack.Push(2)
			case P3:
				vm.operandstack.Push(3)	
			case P4:
				vm.operandstack.Push(4)	
			case P5:
				vm.operandstack.Push(5)
			case P6:
				vm.operandstack.Push(6)
			case P7:
				vm.operandstack.Push(7)
			case P8:
				vm.operandstack.Push(8)
			case P9:
				vm.operandstack.Push(9)
			case ADD:
				S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
				vm.operandstack.Push(S1 + S0)
			case MUL:
				S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
				vm.operandstack.Push(S1 * S0)
			case SUB:
				S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
				vm.operandstack.Push(S1 - S0)
			case DIV:
				S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
				vm.operandstack.Push(S1 / S0)
			case CMP:
				S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
				switch {
					case S0 > S1:
						vm.operandstack.Push(-1)
					case S0 < S1:
						vm.operandstack.Push(1)
					default:
						vm.operandstack.Push(0)
				}
			case JMP:
				S0 := vm.operandstack.Pop()
				t := vm.eip + S0
				if t < 0 || t >= len(code) {
					panic("jump out of code range")
				}
				vm.eip += S0
			case JNE:
				S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
				if S1 == 0 {
					t := vm.eip + S0
					if t < 0 || t > len(code) {
						panic("jump out of code range")
					}
					vm.eip += S0
				}
			case CALL:
				vm.callstack.Push(vm.eip)
				vm.eip = vm.operandstack.Pop()
			case RET:
				vm.eip = vm.callstack.Pop()
			case PEEK:
				vm.operandstack.Push(vm.memory.Get(vm.operandstack.Pop()))
			case POKE:
				vm.memory.Set(vm.operandstack.Pop(), vm.operandstack.Pop())
			case PICK:
				vm.operandstack.Push(vm.operandstack.Get(
					vm.operandstack.Sp - vm.operandstack.Pop() - 1))
			case ROLL:
				S0 := vm.operandstack.Pop()
				t := vm.operandstack.Sp - S0 - 1
				v := vm.operandstack.Get(t)
				vm.operandstack.Del(t)
				vm.operandstack.Push(v)
			case DROP:
				vm.operandstack.Pop()
			case END:
				vm.eip = len(code)
			default:
				panic("unknown bytecode")
		}
		
		if vm.logoperandstack {
			log = append(log, fmt.Sprintf("%c - %v", 
				code[oldeip], 
				vm.operandstack.ValidValues()))
		}
	}
}