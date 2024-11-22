// Package gstream
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/11
package gstream

import (
	"sort"

	"github.com/hyphennn/glambda/gutils"
)

type SliceStream[T any] struct {
	s []T
}

func AsSliceStream[T any](s []T) *SliceStream[T] {
	return &SliceStream[T]{s}
}

func MapAsSliceStream[K comparable, V, T any, M ~map[K]V](m M, fc func(K, V) T) *SliceStream[T] {
	return ToSliceStream(
		AsMapStream(m),
		func(k K, v V) T {
			return fc(k, v)
		},
	)
}

func ToMapStream[K comparable, T, V any](s *SliceStream[T], fc func(T) (K, V)) *MapStream[K, V] {
	m := make(map[K]V, len(s.s))
	s.ForEach(func(t T) {
		k, v := fc(t)
		m[k] = v
	})
	return AsMapStream(m)
}

func ToOtherSliceStream[T1, T2 any](s *SliceStream[T1], fc func(T1) T2) *SliceStream[T2] {
	s2 := make([]T2, 0, len(s.s))
	s.ForEach(func(t1 T1) {
		s2 = append(s2, fc(t1))
	})
	return AsSliceStream(s2)
}

func (s *SliceStream[T]) ForEach(fc func(T)) *SliceStream[T] {
	for _, v := range s.s {
		fc(v)
	}
	return s
}

func (s *SliceStream[T]) Filter(fc func(T) bool) *SliceStream[T] {
	s2 := make([]T, 0, len(s.s)/2)
	s.ForEach(func(t T) {
		if fc(t) {
			s2 = append(s2, t)
		}
	})
	s.s = s2
	return s
}

// Convert AKA: Map
func (s *SliceStream[T]) Convert(fc func(T) T) *SliceStream[T] {
	s2 := make([]T, 0, len(s.s))
	s.ForEach(func(t T) {
		s2 = append(s2, fc(t))
	})
	s.s = s2
	return s
}

// Map an alias of [Convert]
func (s *SliceStream[T]) Map(fc func(T) T) *SliceStream[T] {
	return s.Convert(fc)
}

func (s *SliceStream[T]) distinct() *SliceStream[T] {
	// todo
	panic("need implement")
}

func (s *SliceStream[T]) Sort(less func(t1, t2 T) bool) *SliceStream[T] {
	sort.Slice(s.s, func(i, j int) bool {
		return less(s.s[i], s.s[j])
	})
	return s
}

func (s *SliceStream[T]) Count() int {
	return len(s.s)
}

func (s *SliceStream[T]) Collect() []T {
	return s.s
}

func (s *SliceStream[T]) CollectNoError() ([]T, error) {
	return gutils.NoError(s.Collect())
}
