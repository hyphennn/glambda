// Package gmap
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package gmap

// Union merges multiple maps into a single map.
// If the same key exists in multiple maps, the value from the last map is used.
//
// EXAMPLE:
//
//	m1 := map[int]string{1: "a", 2: "b"}
//	m2 := map[int]string{2: "c", 3: "d"}
//	Union(m1, m2) => map[int]string{1: "a", 2: "c", 3: "d"}
//
// HINT:
//
//   - Use [UnionOnConflict] if you need custom conflict resolution.
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

// UnionOnConflict merges multiple maps into a single map with custom conflict resolution.
// If the same key exists in multiple maps, the provided function fc is used to resolve the conflict.
//
// EXAMPLE:
//
//	m1 := map[int]string{1: "a", 2: "b"}
//	m2 := map[int]string{2: "c", 3: "d"}
//	fc := func(k int, old, new string) string { return old + new }
//	UnionOnConflict([]map[int]string{m1, m2}, fc) => map[int]string{1: "a", 2: "bc", 3: "d"}
//
// HINT:
//
//   - Use [UseNew], [UseOld], or [UseZero] as predefined conflict resolution strategies.
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

// OnConflict defines a function type for resolving key conflicts during map merging.
//
// EXAMPLE:
//
//	fc := func(k int, old, new string) string { return old + new }
//	UnionOnConflict([]map[int]string{m1, m2}, fc)
type OnConflict[K any, V any] func(k K, old, new V) V

// UseNew resolves conflicts by always using the new value.
//
// EXAMPLE:
//
//	UseNew(1, "old", "new") => "new"
var _ OnConflict[any, any] = UseNew[any, any]

// UseOld resolves conflicts by always using the old value.
//
// EXAMPLE:
//
//	UseOld(1, "old", "new") => "old"
var _ OnConflict[any, any] = UseOld[any, any]

// UseZero resolves conflicts by always using the zero value of the type.
//
// EXAMPLE:
//
//	UseZero(1, "old", "new") => ""
var _ OnConflict[any, any] = UseZero[any, any]

// UseNew resolves conflicts by always using the new value.
//
// EXAMPLE:
//
//	UseNew(1, "old", "new") => "new"
func UseNew[K any, V any](k K, old, new V) V {
	return new
}

// UseOld resolves conflicts by always using the old value.
//
// EXAMPLE:
//
//	UseOld(1, "old", "new") => "old"
func UseOld[K any, V any](k K, old, new V) V {
	return old
}

// UseZero resolves conflicts by always using the zero value of the type.
//
// EXAMPLE:
//
//	UseZero(1, "old", "new") => ""
func UseZero[K any, V any](k K, old, new V) (v V) {
	return
}
