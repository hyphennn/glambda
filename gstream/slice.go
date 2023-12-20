// Package gstream
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/11
package gstream

type SliceStream[T any] struct {
	s []T
}

func AsSliceStream[T any](s []T) *SliceStream[T] {
	return &SliceStream[T]{s}
}

func ToMapStream[K comparable, T, V any](s *SliceStream[T], fc func(T) (K, V)) *MapStream[K, V] {
	m := make(map[K]V, len(s.s))
	s.ForEach(func(t T) {
		k, v := fc(t)
		m[k] = v
	})
	return AsMapStream(m)
}

func (s *SliceStream[T]) ForEach(fc func(T)) {
	for _, v := range s.s {
		fc(v)
	}
}

func (s *SliceStream[T]) Count() int {
	return len(s.s)
}

func (s *SliceStream[T]) Collect() []T {
	return s.s
}
