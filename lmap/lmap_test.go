// Package lmap
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package lmap_test

import (
	"strconv"
	"testing"

	"github.com/hyphennn/glamda/internal/assert"
	"github.com/hyphennn/glamda/lmap"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		map[string]string{"1": "1", "4": "4", "5": "5"},
		lmap.Map(map[int]int{1: 1, 4: 4, 5: 5}, func(k1 int, v1 int) (string, string) {
			return strconv.Itoa(k1), strconv.Itoa(v1)
		}))
}
