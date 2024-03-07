package gmap

import (
	"github.com/hyphennn/glambda/gutils"
)

func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, fc func(K1, V1) (K2, V2)) map[K2]V2 {
	ret := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2 := fc(k1, v1)
		ret[k2] = v2
	}
	return ret
}

func ForEach[K comparable, V any](m map[K]V, fc func(K, V)) {
	for k, v := range m {
		fc(k, v)
	}
}

func Reverse[K, V comparable](m map[K]V) map[V]K {
	ret := make(map[V]K, len(m))
	for k, v := range m {
		ret[v] = k
	}
	return ret
}

func SafeStore[K comparable, V any, M ~map[K]V](m M, k K, v V) {
	if m == nil {
		m = make(map[K]V)
	}
	m[k] = v
}

func ToSlice[K comparable, V, T any](m map[K]V, fc KVTrans[K, V, T]) []T {
	ret := make([]T, 0, len(m))
	for k, v := range m {
		ret = append(ret, fc(k, v))
	}
	return ret
}

type KVTrans[K comparable, V, T any] func(K, V) T

var _ KVTrans[any, any, any] = UseKey[any, any]
var _ KVTrans[any, any, any] = UseValue[any, any]
var _ KVTrans[any, any, *gutils.Pair[any, any]] = UsePair[any, any]

func UseKey[K comparable, V any](k K, v V) K {
	return k
}

func UseValue[K comparable, V any](k K, v V) V {
	return v
}

func UsePair[K comparable, V any](k K, v V) *gutils.Pair[K, V] {
	return gutils.MakePair(k, v)
}

func CollectKey[K comparable, V any](m map[K]V) []K {
	return ToSlice(m, UseKey[K, V])
}

func CollectValue[K comparable, V any](m map[K]V) []V {
	return ToSlice(m, UseValue[K, V])
}
