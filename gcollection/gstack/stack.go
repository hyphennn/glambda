// Package gstack
// Author: hyphen
// Copyright 2024 hyphen. All rights reserved.
// Create-time: 2024/1/3
package gstack

import (
	"github.com/hyphennn/glamda/internal"
)

type node[T any] struct {
	val  T
	next *node[T]
}

type Stack[T any] struct {
	top  *node[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Push(t T) {
	s.top = &node[T]{t, s.top}
	s.size++
}

func (s *Stack[T]) PushN(ts ...T) {
	for _, t := range ts {
		s.Push(t)
	}
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		return internal.Zero[T](), false
	}
	ot := s.top
	s.top = s.top.next
	s.size--
	return ot.val, true
}

func (s *Stack[T]) PopN(n int) []T {
	ret := []T{}
	for i := 0; i < n; i++ {
		v, ok := s.Pop()
		if !ok {
			break
		}
		ret = append(ret, v)
	}
	return ret
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.top == nil {
		return internal.Zero[T](), false
	}
	return s.top.val, true
}

func (s *Stack[T]) PeekN(n int) []T {
	ptr := s.top
	ret := []T{}
	for i := 0; i < n; i++ {
		if ptr == nil {
			break
		}
		ret = append(ret, ptr.val)
		ptr = ptr.next
	}
	return ret
}
