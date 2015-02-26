package gohackvm

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MEMORY_SIZE       = 16384
	OPERANDSTACK_SIZE = 256
	CALLSTACK_SIZE    = 256
)

type HackVM struct {
	eip             int
	memory          *Ram
	operandstack    *Stack
	callstack       *Stack
	logoperandstack bool
}

func NewHackVM(logoperandstack bool) *HackVM {
	return &HackVM{0,
		NewRam(MEMORY_SIZE),
		NewStack(OPERANDSTACK_SIZE),
		NewStack(CALLSTACK_SIZE),
		logoperandstack}
}

func (vm *HackVM) SetInitialMemory(initmems string) {
	if initmems != "" {
		initmem := strings.Split(initmems, ",")
		for i, s := range initmem {
			m, err := strconv.Atoi(s)
			if err == nil {
				vm.memory.Set(i, m)
			}
		}
	}
}

func (vm *HackVM) RunProgramString(code string) {
	vm.RunProgram([]byte(code))
}

func (vm *HackVM) RunProgram(code []byte) {
	var log *bytes.Buffer
	if vm.logoperandstack {
		log = bytes.NewBufferString("\n")
	}
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf(
				"%s caused by %c with eip %d\noperandstack=%d, callstack=%d\n",
				e.(string),
				code[vm.eip],
				vm.eip,
				vm.operandstack.Sp,
				vm.callstack.Sp)
		}
		if vm.logoperandstack {
			log.WriteTo(os.Stdout)
		}
	}()
	for vm.eip >= 0 && vm.eip < len(code) {
		neip := vm.eip + 1
		switch code[vm.eip] {
		case ' ':
		case '\n':
		case 'p':
			print(vm.operandstack.Pop())
		case 'P':
			print(string(vm.operandstack.Pop()))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			vm.operandstack.Push(int(byte(code[vm.eip]) - '0'))
		case '+':
			S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
			vm.operandstack.Push(S1 + S0)
		case '*':
			S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
			vm.operandstack.Push(S1 * S0)
		case '-':
			S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
			vm.operandstack.Push(S1 - S0)
		case '/':
			S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
			vm.operandstack.Push(S1 / S0)
		case ':':
			S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
			switch {
			case S0 > S1:
				vm.operandstack.Push(-1)
			case S0 < S1:
				vm.operandstack.Push(1)
			default:
				vm.operandstack.Push(0)
			}
		case 'g':
			S0 := vm.operandstack.Pop()
			t := neip + S0
			if t < 0 || t >= len(code) {
				panic("jump out of code range")
			}
			neip = t
		case '?':
			S0, S1 := vm.operandstack.Pop(), vm.operandstack.Pop()
			if S1 == 0 {
				t := neip + S0
				if t < 0 || t > len(code) {
					panic("jump out of code range")
				}
				neip = t
			}
		case 'c':
			vm.callstack.Push(neip)
			neip = vm.operandstack.Pop()
		case '$':
			neip = vm.callstack.Pop()
		case '<':
			vm.operandstack.Push(vm.memory.Get(vm.operandstack.Pop()))
		case '>':
			vm.memory.Set(vm.operandstack.Pop(), vm.operandstack.Pop())
		case '^':
			vm.operandstack.Push(vm.operandstack.Get(
				vm.operandstack.Sp - vm.operandstack.Pop() - 1))
		case 'v':
			S0 := vm.operandstack.Pop()
			t := vm.operandstack.Sp - S0 - 1
			v := vm.operandstack.Get(t)
			vm.operandstack.Del(t)
			vm.operandstack.Push(v)
		case 'd':
			vm.operandstack.Pop()
		case '!':
			neip = len(code)
		default:
			panic("unknown bytecode")
		}
		if vm.logoperandstack {
			fmt.Fprintf(log, "%d:\t%c\t%v\n",
				vm.eip,
				code[vm.eip],
				vm.operandstack.ValidValues())
		}
		vm.eip = neip
	}
}
