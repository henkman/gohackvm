include $(GOROOT)/src/Make.inc

TARG=gohackvm
GOFILES=\
	main.go\
	gohackvm.go\
	bytecodes.go\
	ram.go\
	stack.go\
	
include $(GOROOT)/src/Make.cmd