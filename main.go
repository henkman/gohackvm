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

var _filename string
var _logstack bool
var _code string

func init() {
	flag.BoolVar(&_logstack, "l", false, "shall the stack be logged")
	flag.StringVar(&_filename, "f", "", "file to be executed")
	flag.StringVar(&_code, "c", "", "code to be executed")
}

func main() {
	flag.Parse()
	
	switch {
		case _code != "":
			RunProgramString(_code, _logstack)
			
		case _filename != "":
			code, err := ioutil.ReadFile(_filename)
			if err != nil {
				println(err.String())
			}
			
			RunProgram(code, _logstack)
			
		default:
			flag.Usage()
	}
}