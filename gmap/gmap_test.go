// Package gmap
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package gmap_test

import (
	"strconv"
	"testing"

	"github.com/hyphennn/glambda/gmap"
	"github.com/hyphennn/glambda/internal/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		map[string]string{"1": "1", "4": "4", "5": "5"},
		gmap.Map(map[int]int{1: 1, 4: 4, 5: 5}, func(k1 int, v1 int) (string, string) {
			return strconv.Itoa(k1), strconv.Itoa(v1)
		}))
}

func TestUnion(t *testing.T) {
	assert.Equal(t,
		map[string]string{"1": "1", "4": "4", "5": "5"},
		gmap.Union(map[string]string{"1": "1", "4": "4"}, map[string]string{"4": "4", "5": "5"}),
	)
}
