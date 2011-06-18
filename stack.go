package main

type Stack struct {
	Values []int
	Sp int
}

func NewStack(maxValues int) *Stack {
	return &Stack{make([]int, maxValues), 0}
}

func (s *Stack) Pop() int {
	if s.Sp == 0 {
		panic("stack underflow")
	}
	
	s.Sp--
	return s.Values[s.Sp]
}

func (s *Stack) Push(v int) {
	if s.Sp == len(s.Values) {
		panic("stack overflow")
	}
	
	s.Values[s.Sp] = v
	s.Sp++
}

func (s *Stack) Get(index int) int {
	if index > s.Sp || index < 0 {
		panic("index out of stack range")
	}
	
	return s.Values[index]
}

func (s *Stack) Set(index int, v int) {
	if index > s.Sp || index < 0 {
		panic("index out of stack range")
	}
	
	s.Values[index] = v
}

func (s *Stack) Del(index int) {
	if index > s.Sp || index < 0 {
		panic("index out of stack range")
	}
	
	if s.Sp == index || s.Sp == 1 {
		s.Sp--
		return
	}
	
	for i, _ := range(s.Values[index:s.Sp]) {
		s.Values[index + i] = s.Values[index + i + 1]
	}
	
	s.Sp--
}

func (s *Stack) ValidValues() []int {
	return s.Values[:s.Sp]
}