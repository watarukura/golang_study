package main

import "fmt"

type Stack struct {
	store []string
	limit int
}

func (s *Stack) Pop() string {
	if len(s.store) == 0 {
		return ""
	}
	var x string
	x, s.store = s.store[len(s.store)-1], s.store[:len(s.store)-1]
	return x
}
func (s *Stack) Push(ss string) {
	s.store = append(s.store, ss)
	if len(s.store) > s.limit {
		s.store = s.store[1:]
	}
}

func main() {
	s := &Stack{}
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Pop()

	fmt.Println(s.Pop())
	fmt.Println(s.store)
}
