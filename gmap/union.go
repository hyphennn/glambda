// Package gmap
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package gmap

func Union[K comparable, V any](ms ...map[K]V) map[K]V {
	if len(ms) == 0 {
		return make(map[K]V)
	}
	if len(ms) == 1 {
		return ms[0]
	}
	l := 0
	for _, m := range ms {
		l += len(m)
	}
	ret := make(map[K]V, l)

	if l == 0 {
		return ret
	}

	for _, m := range ms {
		for k, v := range m {
			ret[k] = v
		}
	}
	return ret
}

func UnionOnConflict[K comparable, V any, M ~map[K]V](ms []M, fc OnConflict[K, V]) map[K]V {
	if len(ms) == 0 {
		return make(map[K]V)
	}
	if len(ms) == 1 {
		return ms[0]
	}

	l := 0
	for _, m := range ms {
		l += len(m)
	}
	ret := make(map[K]V, l)

	if l == 0 {
		return ret
	}

	for _, m := range ms {
		for k, v := range m {
			if v0, ok := ret[k]; ok {
				ret[k] = fc(k, v0, v)
				continue
			}
			ret[k] = v
		}
	}
	return ret
}

type OnConflict[K comparable, V any] func(k K, old, new V) V

var _ OnConflict[any, any] = UseNew[any, any]
var _ OnConflict[any, any] = UseOld[any, any]
var _ OnConflict[any, any] = UseZero[any, any]

func UseNew[K comparable, V any](k K, old, new V) V {
	return new
}

func UseOld[K comparable, V any](k K, old, new V) V {
	return old
}

func UseZero[K comparable, V any](k K, old, new V) (v V) {
	return
}
