package main

//PtrStack is a simple int stack,
//intended to be used as a stack of slice/array pointers
type PtrStack struct {
	Ptrs []int
}

//NewPtrStack retruns a new PtrStack
func NewPtrStack() *PtrStack {
	return new(PtrStack)
}

//IsEmpty returns true if the given PtrStack is empty.
func (s *PtrStack) IsEmpty() bool {
	return len(s.Ptrs) == 0
}

//Push pushes i onto the given PtrStack.
func (s *PtrStack) Push(i int) {
	s.Ptrs = append(s.Ptrs, i)
}

//Peek returns the value ontop of the given PtrStack
func (s *PtrStack) Peek() int {
	return s.Ptrs[len(s.Ptrs)-1]
}

//Pop returns the value ontop of the given PtrStack and pops the top value.
func (s *PtrStack) Pop() (i int) {
	i = s.Peek()
	s.Ptrs = s.Ptrs[:len(s.Ptrs)-1]
	return
}
