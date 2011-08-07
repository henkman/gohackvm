package main

import (
	"flag"
	"io/ioutil"
)

const (
	MEMORY_SIZE = 16384
	OPERANDSTACK_SIZE = 256
	CALLSTACK_SIZE = 256
)

var _logstack bool
var _code string
var _codefile string
var _mem string
var _memfile string

func init() {
	flag.BoolVar(&_logstack, "l", false, "shall the stack be logged")
	flag.StringVar(&_code, "c", "", "code to be executed")
	flag.StringVar(&_codefile, "cf", "", "file containing code to be executed")
	flag.StringVar(&_mem, "m", "", "the initial memory (cell0,cell1,...)")
	flag.StringVar(&_memfile, "mf", "", "file containing initial memory (cell0,cell1,...)")
}

func main() {
	flag.Parse()
	
	if _memfile != "" {
		m, err := ioutil.ReadFile(_memfile)
		if err != nil {
			println(err.String())
			return
		}
		
		_mem = string(m)
	}
	
	switch {
		case _code != "":
			RunProgramString(_code, _logstack, _mem)
			
		case _codefile != "":
			code, err := ioutil.ReadFile(_codefile)
			if err != nil {
				println(err.String())
				return
			}
			
			RunProgram(code, _logstack, _mem)
			
		default:
			flag.Usage()
	}
}