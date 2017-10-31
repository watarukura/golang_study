package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	ls := &Stack{limit: 2}
	ls.Push("dataA")
	ls.Push("dataB")
	ls.Push("dataC")


	v := ls.Pop()
	if v != "dataC" {
		t.Error("Error!")
	}
	v = ls.Pop()
	if v != "dataB" {
		t.Error("Error!")
	}

	v = ls.Pop()
	if v != "" {
		t.Error("Error!")
	}
}
