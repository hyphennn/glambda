// Package lslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package lslice

func Map[F, T any](s []F, f func(F) T) []T {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		ret = append(ret, f(v))
	}
	return ret
}

func ToMap[F, V any, K comparable](s []F, f func(F) (K, V)) map[K]V {
	ret := make(map[K]V, len(s))
	for _, e := range s {
		k, v := f(e)
		ret[k] = v
	}
	return ret
}

func TryMap[F, T any](s []F, f func(F) (T, error)) ([]T, error) {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		t, e := f(v)
		if e != nil {
			return ret, e
		}
		ret = append(ret, t)
	}
	return ret, nil
}

func Filter[F any](s []F, f func(F) bool) []F {
	ret := make([]F, 0, len(s)/2)
	for _, v := range s {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
