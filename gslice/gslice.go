// Package gslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gslice

import (
	"github.com/hyphennn/glambda/internal"
)

func Map[F, T any](s []F, fc func(F) T) []T {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		ret = append(ret, fc(v))
	}
	return ret
}

func ToMap[F, V any, K comparable](s []F, fc func(F) (K, V)) map[K]V {
	ret := make(map[K]V, len(s))
	for _, e := range s {
		k, v := fc(e)
		ret[k] = v
	}
	return ret
}

func TryMap[F, T any](s []F, fc func(F) (T, error)) ([]T, error) {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		t, e := fc(v)
		if e != nil {
			return ret, e
		}
		ret = append(ret, t)
	}
	return ret, nil
}

func Filter[F any](s []F, fc func(F) bool) []F {
	ret := make([]F, 0, len(s)/2)
	for _, v := range s {
		if fc(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Reject[T any](s []T, fc func(T) bool) []T {
	ret := make([]T, 0, len(s)/2)
	for _, v := range s {
		if !fc(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func ForEach[T any](s []T, fc func(T)) {
	for _, v := range s {
		fc(v)
	}
}

func ForEachIdx[T any](s []T, fc func(int, T)) {
	for i, v := range s {
		fc(i, v)
	}
}

func Find[T any](s []T, f func(T) bool) (T, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	return internal.Zero[T](), false
}

func GroupBy[K comparable, T any, S ~[]T](s S, f func(T) K) map[K]S {
	m := make(map[K]S)
	for i := range s {
		k := f(s[i])
		m[k] = append(m[k], s[i])
	}
	return m
}

func Contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if v == vv {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](s []T, vs ...T) bool {
	m := make(map[T]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	for _, v := range s {
		delete(m, v)
		if len(m) == 0 {
			return true
		}
	}
	return len(m) == 0
}

func ContainsAny[T comparable](s []T, vs ...T) bool {
	m := make(map[T]struct{}, len(vs))
	for _, v := range vs {
		m[v] = struct{}{}
	}
	for _, v := range s {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}
