// Package gstream
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/11
package gstream

import (
	"github.com/hyphennn/glambda/gutils"
)

type MapStream[K comparable, V any] struct {
	m map[K]V
}

func AsMapStream[K comparable, V any, M ~map[K]V](m M) *MapStream[K, V] {
	return &MapStream[K, V]{m}
}

func SliceAsMapStream[K comparable, V, T any](s []T, fc func(T) (K, V)) *MapStream[K, V] {
	return ToMapStream(AsSliceStream(s), func(t T) (K, V) {
		return fc(t)
	})
}

func ToSliceStream[K comparable, T, V any](m *MapStream[K, V], fc func(K, V) T) *SliceStream[T] {
	s := make([]T, 0, len(m.m))
	m.ForEach(func(k K, v V) {
		s = append(s, fc(k, v))
	})
	return AsSliceStream(s)
}

func ToOtherMapStream[K1, K2 comparable, V1, V2 any](m *MapStream[K1, V1], fc func(K1, V1) (K2, V2)) *MapStream[K2, V2] {
	m2 := make(map[K2]V2, len(m.m))
	m.ForEach(func(k1 K1, v1 V1) {
		k2, v2 := fc(k1, v1)
		m2[k2] = v2
	})
	return AsMapStream(m2)
}

func (m *MapStream[K, V]) ForEach(fc func(K, V)) *MapStream[K, V] {
	for k, v := range m.m {
		fc(k, v)
	}
	return m
}

func (m *MapStream[K, V]) Filter(fc func(K, V) bool) *MapStream[K, V] {
	for k, v := range m.m {
		if !fc(k, v) {
			// todo benchmark一下delete和新建map的速度
			delete(m.m, k)
		}
	}
	return m
}

func (m *MapStream[K, V]) Limit(n int) *MapStream[K, V] {
	if n >= len(m.m) {
		return m
	}
	m1, idx := make(map[K]V, n), 0
	for k, v := range m.m {
		m1[k] = v
		idx++
		if idx > n {
			break
		}
	}
	return m
}

func (m *MapStream[K, V]) Convert(fc func(K, V) (K, V)) *MapStream[K, V] {
	m2 := make(map[K]V, len(m.m))
	for k, v := range m.m {
		k2, v2 := fc(k, v)
		m2[k2] = v2
	}
	m.m = m2
	return m
}

func (m *MapStream[K, V]) Collect() map[K]V {
	return m.m
}

func (m *MapStream[K, V]) CollectNoError() (map[K]V, error) {
	return gutils.NoError(m.Collect())
}
