// Package gmap
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package gmap_test

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/hyphennn/glambda/gmap"
	"github.com/hyphennn/glambda/gutils"
	"github.com/hyphennn/glambda/internal/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		map[string]string{"1": "1", "4": "4", "5": "5"},
		gmap.Map(map[int]int{1: 1, 4: 4, 5: 5}, func(k1 int, v1 int) (string, string) {
			return strconv.Itoa(k1), strconv.Itoa(v1)
		}))
}

func TestForEach(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	var result string
	gmap.ForEach(m, func(k int, v string) {
		result += fmt.Sprintf("%d:%s ", k, v)
	})
	assert.Equal(t, "1:a 2:b ", result)
}

func TestReverse(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	reversed := gmap.Reverse(m)
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, reversed)
}

func TestSafeStore(t *testing.T) {
	var m map[int]string
	m = gmap.SafeStore(m, 1, "a")
	assert.Equal(t, map[int]string{1: "a"}, m)
}

func TestToSlice(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	result := gmap.ToSlice(m, func(k int, v string) string {
		return fmt.Sprintf("%d:%s", k, v)
	})
	sort.Strings(result)
	assert.Equal(t, []string{"1:a", "2:b"}, result)
}

func TestUseKey(t *testing.T) {
	assert.Equal(t, 1, gmap.UseKey(1, "a"))
}

func TestUseValue(t *testing.T) {
	assert.Equal(t, "a", gmap.UseValue(1, "a"))
}

func TestUsePair(t *testing.T) {
	pair := gmap.UsePair(1, "a")
	assert.Equal(t, gutils.MakePair(1, "a"), pair)
}

func TestCollectKey(t *testing.T) {
	assert.Equal(t,
		3,
		len(gmap.CollectKey(map[int]int{1: 1, 2: 2, 3: 3})),
	)
}

func TestCollectValue(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	values := gmap.CollectValue(m)
	assert.Equal(t, 3, len(values))
}

func TestContainsAll(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	assert.True(t, gmap.ContainsAll(m, 1, 2))
	assert.False(t, gmap.ContainsAll(m, 1, 3))
}

func TestContainsAny(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	assert.True(t, gmap.ContainsAny(m, 1, 3))
	assert.False(t, gmap.ContainsAny(m, 3, 4))
}

func TestContainsMapAll(t *testing.T) {
	parent := map[int]string{1: "a", 2: "b"}
	child := map[int]string{1: "a"}
	assert.True(t, gmap.ContainsMapAll(parent, child))
	assert.False(t, gmap.ContainsMapAll(parent, map[int]string{1: "c"}))
}

func TestContainsMapAny(t *testing.T) {
	parent := map[int]string{1: "a", 2: "b"}
	child := map[int]string{1: "a"}
	assert.True(t, gmap.ContainsMapAny(parent, child))
	assert.False(t, gmap.ContainsMapAny(parent, map[int]string{1: "c"}))
}

func TestClone(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	cloned := gmap.Clone(m)
	assert.Equal(t, m, cloned)
}
