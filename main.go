package main

import (
	"flag"
	"github.com/papplampe/gohackvm/gohackvm"
	"io/ioutil"
	"log"
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
			log.Fatal(err)
		}
		_mem = string(m)
	}
	switch {
	case _code != "":
		vm := gohackvm.NewHackVM(_logstack)
		vm.SetInitialMemory(_mem)
		vm.RunProgramString(_code)
	case _codefile != "":
		code, err := ioutil.ReadFile(_codefile)
		if err != nil {
			log.Fatal(err)
		}
		vm := gohackvm.NewHackVM(_logstack)
		vm.SetInitialMemory(_mem)
		vm.RunProgram(code)
	default:
		flag.Usage()
	}
}
