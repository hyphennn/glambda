package gmap

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

func ToSlice[K comparable, V, T any](m map[K]V, fc func(K, V) T) []T {
	ret := make([]T, 0, len(m))
	for k, v := range m {
		ret = append(ret, fc(k, v))
	}
	return ret
}

func CollectKey[K comparable, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func CollectValue[K comparable, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
