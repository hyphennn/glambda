// Package lslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package lslice_test

import (
	"strconv"
	"testing"

	"github.com/hyphennn/glamda/internal/assert"
	"github.com/hyphennn/glamda/lslice"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		lslice.Map([]int{1, 1, 4, 5, 1, 4}, strconv.Itoa),
		[]string{"1", "1", "4", "5", "1", "4"},
	)
}

func TestToMap(t *testing.T) {
	assert.Equal(t,
		map[int]bool{1: true, 4: true, 5: true},
		lslice.ToMap([]int{1, 1, 4, 5, 1, 4}, func(f int) (int, bool) { return f, true }),
	)
}

func TestTryMap(t *testing.T) {
	m, err := lslice.TryMap([]string{"1", "1", "4", "5", "1", "4"}, strconv.Atoi)
	assert.Nil(t, err)
	assert.Equal(t, m, []int{1, 1, 4, 5, 1, 4})

	m2, err := lslice.TryMap([]string{"1", "1", "4", "5a", "1", "4"}, strconv.Atoi)
	t.Log(err)
	assert.NotNil(t, err)
	assert.Equal(t, m2, []int{1, 1, 4})
}

func TestFilter(t *testing.T) {
	assert.Equal(t,
		lslice.Filter([]int{1, 1, 4, 5, 1, 4}, func(i int) bool {
			return i%2 == 1
		}),
		[]int{1, 1, 5, 1},
	)
}
