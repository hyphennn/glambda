// Package gvalue
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package gvalue

import (
	"github.com/hyphennn/glambda/internal/constraints"
)

func Sum[T constraints.Addable](s ...T) T {
	var ret T
	for _, v := range s {
		ret += v
	}
	return ret
}

func Max[T constraints.Ordered](s0 T, s ...T) T {
	ret := s0
	for _, v := range s {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func Min[T constraints.Ordered](s0 T, s ...T) T {
	ret := s0
	for _, v := range s {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func Zero[T any]() (t T) {
	return
}
