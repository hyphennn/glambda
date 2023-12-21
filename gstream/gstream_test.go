// Package gstream
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/11
package gstream

import (
	"fmt"
	"testing"
)

func Test0(t *testing.T) {
	s := []int{1, 1, 4, 5, 1, 4, 114514, 0, 0, 0}
	s2 := AsSliceStream(s).Filter(func(i int) bool {
		return i%2 == 0
	}).Convert(func(i int) int {
		return i + 1
	}).Sort(func(t1, t2 int) bool {
		return t1 < t2
	}).Collect()
	fmt.Println(s2)
}
