// Package gstream
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/11
package gstream

type MapStream[K comparable, V any] struct {
	m map[K]V
	s *SliceStream[K]
}

func AsMapStream[K comparable, V any, M ~map[K]V](m M) *MapStream[K, V] {
	return &MapStream[K, V]{m, nil}
}

func ToSliceStream[K comparable, T, V any](m *MapStream[K, V], fc func(K, V) T) *SliceStream[T] {
	s := make([]T, 0, len(m.m))
	m.ForEach(func(k K, v V) {
		s = append(s, fc(k, v))
	})
	return AsSliceStream(s)
}

func (m *MapStream[K, V]) ForEach(fc func(K, V)) *MapStream[K, V] {
	for k, v := range m.m {
		fc(k, v)
	}
	return m
}

func (m *MapStream[K, V]) Collect() map[K]V {
	return m.m
}
