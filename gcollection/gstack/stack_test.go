// Package gstack
// Author: hyphen
// Copyright 2024 hyphen. All rights reserved.
// Create-time: 2024/1/3
package gstack

import (
	"testing"

	"github.com/hyphennn/glamda/internal/assert"
)

func TestStack(t *testing.T) {
	stk := NewStack[int]()
	stk.PushN([]int{1, 1, 4, 5, 1, 4}...)
	stk.Push(1)
	v, ok := stk.Peek()
	vs := stk.PeekN(5)
	assert.Equal(t, v, 1)
	assert.True(t, ok)
	assert.Equal(t, stk.Size(), 7)
	assert.Equal(t, vs, []int{1, 4, 1, 5, 4})
	v, ok = stk.Pop()
	assert.Equal(t, v, 1)
	assert.True(t, ok)
	assert.Equal(t, stk.Size(), 6)
	vs = stk.PopN(5)
	assert.Equal(t, vs, []int{4, 1, 5, 4, 1})
	assert.Equal(t, stk.Size(), 1)
	assert.Equal(t, stk.PeekN(3), []int{1})
	v, ok = stk.Pop()
	v, ok = stk.Pop()
	assert.Equal(t, v, 0)
	assert.False(t, ok)
	stk.Push(1)
	v, ok = stk.Pop()
	assert.Equal(t, v, 1)
	assert.True(t, ok)
}
