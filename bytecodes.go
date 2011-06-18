package main

type Bytecode byte
const (
	NOP Bytecode = ' '  // Do Nothing
	NOP2 Bytecode = '\n' // Do Nothing
	PINT Bytecode = 'p' // Print S0 interpreted as an integer
	PCHR Bytecode = 'P' // Print S0 interpreted as an ASCII character (only the least significant 7 bits of the value are used)
	P0 Bytecode = '0'   // Push the value 0 on the stack
	P1 Bytecode = '1'   // Push the value 1 on the stack
	P2 Bytecode = '2'   // Push the value 2 on the stack
	P3 Bytecode = '3'   // Push the value 3 on the stack
	P4 Bytecode = '4'   // Push the value 4 on the stack
	P5 Bytecode = '5'   // Push the value 5 on the stack
	P6 Bytecode = '6'   // Push the value 6 on the stack
	P7 Bytecode = '7'   // Push the value 7 on the stack
	P8 Bytecode = '8'   // Push the value 8 on the stack
	P9 Bytecode = '9'   // Push the value 9 on the stack
	ADD Bytecode = '+'  // Push S1+S0
	MUL Bytecode = '*'  // Push S1*S0
	SUB Bytecode = '-'  // Push S1-S0
	DIV Bytecode = '/'  // Push S1/S0
	CMP Bytecode = ':'  // Push -1 if S1<S0, 0 if S1=S0, or 1 S1>S0
	JMP Bytecode = 'g'  // Add S0 to the program counter
	JNE Bytecode = '?'  // Add S0 to the program counter if S1 is 0
	CALL Bytecode = 'c' // Push the program counter on the call stack and set the program counter to S0
	RET Bytecode = '$' // Set the program counter to the value pop'ed from the call stack
	PEEK Bytecode = '<'  // Push the value of memory cell S0
	POKE Bytecode = '>'  // Store S1 into memory cell S0
	PICK Bytecode = '^'  // Push a copy of S<S0+1> (ex: 0^ duplicates S0)
	ROLL Bytecode = 'v'  // Remove S<S0+1> from the stack and push it on top (ex: 1v swaps S0 and S1)
	DROP Bytecode = 'd'  // Drop S0
	END Bytecode = '!'  // Terminate the program
)