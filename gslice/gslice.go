// Package gslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gslice

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
