// Package gslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gslice

import (
	"github.com/hyphennn/glambda/gutils"
	"github.com/hyphennn/glambda/gvalue"
)

func Map[F, T any](s []F, fc func(F) T) []T {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		ret = append(ret, fc(v))
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

func ToMap[F, V any, K comparable](s []F, fc func(F) (K, V)) map[K]V {
	ret := make(map[K]V, len(s))
	for _, e := range s {
		k, v := fc(e)
		ret[k] = v
	}
	return ret
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

func FilterMap[F, T any](s []F, fc func(F) (T, bool)) []T {
	ret := make([]T, 0, len(s)/2)
	for _, v := range s {
		if t, ok := fc(v); ok {
			ret = append(ret, t)
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

func Reduce[T any](s []T, fc func(T, T) T) (T, bool) {
	if len(s) == 0 {
		return gvalue.Zero[T](), false
	}
	ret := s[0]
	for _, v := range s[1:] {
		ret = fc(ret, v)
	}
	return ret, true
}

func Fold[T1, T2 any](s []T1, fc func(T2, T1) T2, init T2) T2 {
	ret := init
	for _, v := range s[1:] {
		ret = fc(ret, v)
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
	return gvalue.Zero[T](), false
}

func FindRev[T any](s []T, f func(T) bool) (T, bool) {
	for i := len(s); i >= 0; i-- {
		if f(s[i]) {
			return v, true
		}
	}
	return gvalue.Zero[T](), false
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

func Remove[T comparable](s []T, v T) []T {
	return Filter(s, func(t T) bool {
		return t == v
	})
}

func RemoveN[T comparable](s []T, v T, n int) []T {
	return Filter(s, func(t T) bool {
		if n <= 0 {
			return true
		}
		if t != v {
			return true
		}
		n--
		return false
	})
}

func Distinct[T comparable](s []T) []T {
	return gutils.NewSliceSetFormSlice(s).GetSlice()
}

func DistinctBy[K comparable, V any](s []V, fc func(V) K) []V {
	ss := gutils.NewSliceSet[K, V]()
	for _, v := range s {
		ss.Upsert(fc(v), v)
	}
	return ss.GetSlice()
}

func DeepCopy[T any, S ~[]T](s S) S {
	if s == nil {
		return nil
	}
	return Map(s, func(t T) T { return t })
}
