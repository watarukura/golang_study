package main

import "fmt"

type Stack struct {
	S []string
}

func (s *Stack) Pop() string {
	var x string
	x, s.S = s.S[len(s.S)-1], s.S[:len(s.S)-1]
	return x
}
func (s *Stack) Push(ss string) {
	s.S = append(s.S, ss)
}

func main()  {
	stack := &Stack{}
	ss := "a"
	stack.S = append(stack.S, ss)
	stack.S = append(stack.S, ss)
	stack.Push("b")
	fmt.Println(stack.Pop())
	fmt.Println(stack.S)
}
