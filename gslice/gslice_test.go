// Package gslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gslice_test

import (
	"strconv"
	"testing"

	"github.com/hyphennn/glambda/gslice"
	"github.com/hyphennn/glambda/internal/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		gslice.Map([]int{1, 1, 4, 5, 1, 4}, strconv.Itoa),
		[]string{"1", "1", "4", "5", "1", "4"},
	)

	// 测试空切片
	assert.Equal(t,
		gslice.Map([]int{}, strconv.Itoa),
		[]string{},
	)

	// 测试 nil 切片
	assert.Equal(t,
		gslice.Map(nil, strconv.Itoa),
		[]string{},
	)
}

func TestToMap(t *testing.T) {
	assert.Equal(t,
		map[int]bool{1: true, 4: true, 5: true},
		gslice.ToMap([]int{1, 1, 4, 5, 1, 4}, func(f int) (int, bool) { return f, true }),
	)

	// 测试空切片
	assert.Equal(t,
		map[int]bool{},
		gslice.ToMap([]int{}, func(f int) (int, bool) { return f, true }),
	)

	// 测试 nil 切片
	assert.Equal(t,
		map[int]bool{},
		gslice.ToMap(nil, func(f int) (int, bool) { return f, true }),
	)
}

func TestTryMap(t *testing.T) {
	m, err := gslice.TryMap([]string{"1", "1", "4", "5", "1", "4"}, strconv.Atoi)
	assert.Nil(t, err)
	assert.Equal(t, m, []int{1, 1, 4, 5, 1, 4})

	m2, err := gslice.TryMap([]string{"1", "1", "4", "5a", "1", "4"}, strconv.Atoi)
	t.Log(err)
	assert.NotNil(t, err)
	assert.Equal(t, m2, []int{1, 1, 4})

	// 测试空切片
	m3, err := gslice.TryMap([]string{}, strconv.Atoi)
	assert.Nil(t, err)
	assert.Equal(t, m3, []int{})

	// 测试 nil 切片
	m4, err := gslice.TryMap(nil, strconv.Atoi)
	assert.Nil(t, err)
	assert.Equal(t, m4, []int{})
}

func TestFilter(t *testing.T) {
	assert.Equal(t,
		gslice.Filter([]int{1, 1, 4, 5, 1, 4}, func(i int) bool {
			return i%2 == 1
		}),
		[]int{1, 1, 5, 1},
	)

	// 测试空切片
	assert.Equal(t,
		gslice.Filter([]int{}, func(i int) bool { return i%2 == 1 }),
		[]int{},
	)

	// 测试 nil 切片
	assert.Equal(t,
		gslice.Filter(nil, func(i int) bool { return i%2 == 1 }),
		[]int{},
	)
}

func TestAll(t *testing.T) {
	// 所有元素都满足条件
	assert.True(t, gslice.All([]int{2, 4, 6}, func(i int) bool { return i%2 == 0 }))

	// 有元素不满足条件
	assert.False(t, gslice.All([]int{2, 3, 6}, func(i int) bool { return i%2 == 0 }))

	// 测试空切片
	assert.True(t, gslice.All([]int{}, func(i int) bool { return i%2 == 0 }))

	// 测试 nil 切片
	assert.True(t, gslice.All(nil, func(i int) bool { return i%2 == 0 }))
}

func TestAny(t *testing.T) {
	// 有元素满足条件
	assert.True(t, gslice.Any([]int{2, 3, 6}, func(i int) bool { return i%2 != 0 }))

	// 所有元素都不满足条件
	assert.False(t, gslice.Any([]int{2, 4, 6}, func(i int) bool { return i%2 != 0 }))

	// 测试空切片
	assert.False(t, gslice.Any([]int{}, func(i int) bool { return i%2 != 0 }))

	// 测试 nil 切片
	assert.False(t, gslice.Any(nil, func(i int) bool { return i%2 != 0 }))
}

func TestFirst(t *testing.T) {
	// 有元素满足条件
	first, ok := gslice.First([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 })
	assert.True(t, ok)
	assert.Equal(t, first, 2)

	// 没有元素满足条件
	first, ok = gslice.First([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, first, 0)

	// 测试空切片
	first, ok = gslice.First([]int{}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, first, 0)

	// 测试 nil 切片
	first, ok = gslice.First([]int(nil), func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, first, 0)
}

func TestLast(t *testing.T) {
	// 有元素满足条件
	last, ok := gslice.Last([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 })
	assert.True(t, ok)
	assert.Equal(t, last, 2)

	// 没有元素满足条件
	last, ok = gslice.Last([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, last, 0)

	// 测试空切片
	last, ok = gslice.Last([]int{}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, last, 0)

	// 测试 nil 切片
	last, ok = gslice.Last([]int(nil), func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, last, 0)
}

func TestFilterMap(t *testing.T) {
	// 正常情况
	assert.Equal(t,
		gslice.FilterMap([]int{1, 2, 3, 4}, func(i int) (string, bool) {
			if i%2 == 0 {
				return strconv.Itoa(i), true
			}
			return "", false
		}),
		[]string{"2", "4"},
	)

	// 测试空切片
	assert.Equal(t,
		gslice.FilterMap([]int{}, func(i int) (string, bool) {
			if i%2 == 0 {
				return strconv.Itoa(i), true
			}
			return "", false
		}),
		[]string{},
	)

	// 测试 nil 切片
	assert.Equal(t,
		gslice.FilterMap(nil, func(i int) (string, bool) {
			if i%2 == 0 {
				return strconv.Itoa(i), true
			}
			return "", false
		}),
		[]string{},
	)
}

func TestReject(t *testing.T) {
	// 正常情况
	assert.Equal(t,
		gslice.Reject([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }),
		[]int{1, 3},
	)

	// 测试空切片
	assert.Equal(t,
		gslice.Reject([]int{}, func(i int) bool { return i%2 == 0 }),
		[]int{},
	)

	// 测试 nil 切片
	assert.Equal(t,
		gslice.Reject(nil, func(i int) bool { return i%2 == 0 }),
		[]int{},
	)
}

func TestReduce(t *testing.T) {
	// 正常情况
	assert.Equal(t,
		gslice.Reduce([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }),
		10,
	)

	// 测试空切片
	assert.Equal(t,
		gslice.Reduce([]int{}, func(a, b int) int { return a + b }),
		0,
	)
}

func TestFold(t *testing.T) {
	// 正常情况
	assert.Equal(t,
		gslice.Fold([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }, 10),
		20,
	)

	// 测试空切片
	assert.Equal(t,
		gslice.Fold([]int{}, func(a, b int) int { return a + b }, 10),
		10,
	)
}

func TestForEach(t *testing.T) {
	// 由于 ForEach 没有返回值，我们只能测试其副作用
	// 这里通过一个闭包来验证元素是否被遍历
	values := []int{1, 2, 3}
	index := 0
	gslice.ForEach(values, func(i int) {
		assert.Equal(t, i, values[index])
		index++
	})
	assert.Equal(t, index, len(values))

	// 测试空切片
	index = 0
	gslice.ForEach([]int{}, func(i int) {
		index++
	})
	assert.Equal(t, index, 0)

	// 测试 nil 切片
	index = 0
	gslice.ForEach(nil, func(i int) {
		index++
	})
	assert.Equal(t, index, 0)
}

func TestForEachIdx(t *testing.T) {
	// 由于 ForEachIdx 没有返回值，我们只能测试其副作用
	// 这里通过一个闭包来验证元素是否被遍历
	values := []int{1, 2, 3}
	index := 0
	gslice.ForEachIdx(values, func(i, v int) {
		assert.Equal(t, v, values[i])
		index++
	})
	assert.Equal(t, index, len(values))

	// 测试空切片
	index = 0
	gslice.ForEachIdx([]int{}, func(i, v int) {
		index++
	})
	assert.Equal(t, index, 0)

	// 测试 nil 切片
	index = 0
	gslice.ForEachIdx(nil, func(i, v int) {
		index++
	})
	assert.Equal(t, index, 0)
}

func TestFind(t *testing.T) {
	// 有元素满足条件
	find, ok := gslice.Find([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 })
	assert.True(t, ok)
	assert.Equal(t, find, 2)

	// 没有元素满足条件
	find, ok = gslice.Find([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, find, 0)

	// 测试空切片
	find, ok = gslice.Find([]int{}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, find, 0)

	// 测试 nil 切片
	find, ok = gslice.Find(nil, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, find, 0)
}

func TestFindRev(t *testing.T) {
	// 有元素满足条件
	find, ok := gslice.FindRev([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 })
	assert.True(t, ok)
	assert.Equal(t, find, 2)

	// 没有元素满足条件
	find, ok = gslice.FindRev([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, find, 0)

	// 测试空切片
	find, ok = gslice.FindRev([]int{}, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, find, 0)

	// 测试 nil 切片
	find, ok = gslice.FindRev(nil, func(i int) bool { return i%2 == 0 })
	assert.False(t, ok)
	assert.Equal(t, find, 0)
}

func TestGroupBy(t *testing.T) {
	// 正常情况
	assert.Equal(t,
		gslice.GroupBy([]int{1, 2, 3, 4}, func(i int) int { return i % 2 }),
		map[int][]int{0: {2, 4}, 1: {1, 3}},
	)

	// 测试空切片
	assert.Equal(t,
		gslice.GroupBy([]int{}, func(i int) int { return i % 2 }),
		map[int][]int{},
	)

	// 测试 nil 切片
	assert.Equal(t,
		gslice.GroupBy([]int(nil), func(i int) int { return i % 2 }),
		map[int][]int{},
	)
}

func TestContains(t *testing.T) {
	// 包含元素
	assert.True(t, gslice.Contains([]int{1, 2, 3}, 2))

	// 不包含元素
	assert.False(t, gslice.Contains([]int{1, 2, 3}, 4))

	// 测试空切片
	assert.False(t, gslice.Contains([]int{}, 2))

	// 测试 nil 切片
	assert.False(t, gslice.Contains(nil, 2))
}
