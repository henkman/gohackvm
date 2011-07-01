include $(GOROOT)/src/Make.inc

TARG=gohvm
GOFILES=\
	main.go\
	gohackvm.go\
	bytecodes.go\
	ram.go\
	stack.go\
	
include $(GOROOT)/src/Make.cmd