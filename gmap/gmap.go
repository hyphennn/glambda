package gmap

import (
	"github.com/hyphennn/glambda/gutils"
)

// Map applies function fc to each key and value of map m.
// Results of fc are returned as a new map.
//
// EXAMPLE:
//
//	f := func(k, v int) (string, string) { return strconv.Itoa(k), strconv.Itoa(v) }
//	Map(map[int]int{1: 1}, f) => map[string]string{"1": "1"}
//	Map(map[int]int{}, f)     => map[string]string{}
//
// HINT:
//
//   - Use [MapKeys] if you only need to map the keys.
//   - Use [MapValues] if you only need to map the values.
//   - Use [FilterMap] if you also want to ignore keys/values during mapping.
//   - Use [ToSlice] if you want to "map" both key and value to single element
func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, fc func(K1, V1) (K2, V2)) map[K2]V2 {
	ret := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2 := fc(k1, v1)
		ret[k2] = v2
	}
	return ret
}

// ForEach applies function fc to each key and value of map m.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	ForEach(m, func(k int, v string) { fmt.Printf("%d:%s ", k, v) }) => Output: "1:a 2:b "
func ForEach[K comparable, V any](m map[K]V, fc func(K, V)) {
	for k, v := range m {
		fc(k, v)
	}
}

// Reverse swaps keys and values in map m and returns a new map.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	Reverse(m) => map[string]int{"a": 1, "b": 2}
func Reverse[K, V comparable](m map[K]V) map[V]K {
	ret := make(map[V]K, len(m))
	for k, v := range m {
		ret[v] = k
	}
	return ret
}

// SafeStore safely stores a key-value pair in map m.
// If m is nil, it initializes the map.
//
// EXAMPLE:
//
//	var m map[int]string
//	SafeStore(m, 1, "a") => m = map[int]string{1: "a"}
func SafeStore[K comparable, V any, M ~map[K]V](m M, k K, v V) M {
	if m == nil {
		m = make(map[K]V)
	}
	m[k] = v
	return m
}

// ToSlice applies function fc to each key and value of map m and returns a slice of results.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	ToSlice(m, func(k int, v string) string { return fmt.Sprintf("%d:%s", k, v) }) => []string{"1:a", "2:b"}
func ToSlice[K comparable, V, T any](m map[K]V, fc KVTrans[K, V, T]) []T {
	ret := make([]T, 0, len(m))
	for k, v := range m {
		ret = append(ret, fc(k, v))
	}
	return ret
}

type KVTrans[K any, V, T any] func(K, V) T

// UseKey returns the key as the result of the transformation.
//
// EXAMPLE:
//
//	UseKey(1, "a") => 1
func UseKey[K any, V any](k K, v V) K {
	return k
}

// UseValue returns the value as the result of the transformation.
//
// EXAMPLE:
//
//	UseValue(1, "a") => "a"
func UseValue[K any, V any](k K, v V) V {
	return v
}

// UsePair returns a pair of key and value as the result of the transformation.
//
// EXAMPLE:
//
//	UsePair(1, "a") => &gutils.Pair{Key: 1, Value: "a"}
func UsePair[K any, V any](k K, v V) *gutils.Pair[K, V] {
	return gutils.MakePair(k, v)
}

// CollectKey collects all keys from map m into a slice.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	CollectKey(m) => []int{1, 2}
func CollectKey[K comparable, V any](m map[K]V) []K {
	return ToSlice(m, UseKey[K, V])
}

// CollectValue collects all values from map m into a slice.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	CollectValue(m) => []string{"a", "b"}
func CollectValue[K comparable, V any](m map[K]V) []V {
	return ToSlice(m, UseValue[K, V])
}

// ContainsAll checks if all keys ks exist in map m.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	ContainsAll(m, 1, 2) => true
//	ContainsAll(m, 1, 3) => false
func ContainsAll[K comparable, V any](m map[K]V, ks ...K) bool {
	if (m == nil || len(m) == 0) && len(ks) != 0 {
		return false
	}
	for _, k := range ks {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

// ContainsAny checks if any key in ks exists in map m.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	ContainsAny(m, 1, 3) => true
//	ContainsAny(m, 3, 4) => false
func ContainsAny[K comparable, V any](m map[K]V, ks ...K) bool {
	if m == nil || len(m) == 0 {
		return false
	}
	for _, k := range ks {
		if _, ok := m[k]; ok {
			return true
		}
	}
	return false
}

// ContainsMapAll checks if all key-value pairs in child exist in parent.
//
// EXAMPLE:
//
//	parent := map[int]string{1: "a", 2: "b"}
//	child := map[int]string{1: "a"}
//	ContainsMapAll(parent, child) => true
//	ContainsMapAll(parent, map[int]string{1: "c"}) => false
func ContainsMapAll[K, V comparable, M ~map[K]V](parent, child M) bool {
	if len(parent) < len(child) {
		return false
	}
	for k, v := range child {
		if parent[k] != v {
			return false
		}
	}
	return true
}

// ContainsMapAny checks if any key-value pair in child exists in parent.
//
// EXAMPLE:
//
//	parent := map[int]string{1: "a", 2: "b"}
//	child := map[int]string{1: "a"}
//	ContainsMapAny(parent, child) => true
//	ContainsMapAny(parent, map[int]string{1: "c"}) => false
func ContainsMapAny[K, V comparable, M ~map[K]V](parent, child M) bool {
	for k, v := range child {
		if parent[k] == v {
			return true
		}
	}
	return false
}

// Clone creates a shallow copy of map m.
//
// EXAMPLE:
//
//	m := map[int]string{1: "a", 2: "b"}
//	Clone(m) => map[int]string{1: "a", 2: "b"}
func Clone[K comparable, V any, M ~map[K]V](m M) M {
	if m == nil {
		return nil
	}
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}
