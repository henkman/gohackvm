package gohackvm

type Ram struct {
	Memory []int
}

func NewRam(memSize int) *Ram {
	return &Ram{make([]int, memSize)}
}

func (r *Ram) Get(index int) int {
	if index > len(r.Memory) || index < 0 {
		panic("memory read access violation")
	}
	return r.Memory[index]
}

func (r *Ram) Set(index int, v int) {
	if index > len(r.Memory) || index < 0 {
		panic("memory write access violation")
	}
	r.Memory[index] = v
}
