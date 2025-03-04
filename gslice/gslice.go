// Package gslice
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gslice

import (
	"github.com/hyphennn/glambda/gutils"
	"github.com/hyphennn/glambda/gvalue"
)

// Map applies function fc to each element of slice s with type F.
// Results of fc are returned as a new slice with type T.
//
// EXAMPLE:
//
//	Map([]int{1, 2, 3}, strconv.Itoa) => []string{"1", "2", "3"}
//	Map([]int{}, strconv.Itoa)        => []string{}
//	Map(nil, strconv.Itoa)            => []string{}
//
// HINT:
//
//   - Use [FilterMap] if you also want to ignore some element during mapping.
//   - Use [TryMap] if function fc may fail (return (T, error)).
func Map[F, T any](s []F, fc func(F) T) []T {
	ret := make([]T, 0, len(s))
	for _, v := range s {
		ret = append(ret, fc(v))
	}
	return ret
}

// TryMap applies function fc to each element of slice s with type F.
// If fc returns an error for any element, the function stops and returns the error.
// Otherwise, results of fc are returned as a new slice with type T.
//
// EXAMPLE:
//
//	TryMap([]int{1, 2, 3}, func(i int) (string, error) {
//		if i%2 == 0 {
//			return "", errors.New("even number")
//		}
//		return strconv.Itoa(i), nil
//	}) => ([]string{}, error)
//
// HINT:
//
//   - Use [Map] if function fc cannot fail.
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

// ToMap applies function fc to each element of slice s with type F.
// Results of fc are returned as a map with keys of type K and values of type V.
//
// EXAMPLE:
//
//	ToMap([]int{1, 2, 3}, func(i int) (string, int) {
//		return strconv.Itoa(i), i * i
//	}) => map[string]int{"1": 1, "2": 4, "3": 9}
//
// HINT:
//
//   - Ensure that keys returned by fc are unique to avoid overwriting values.
func ToMap[F, V any, K comparable](s []F, fc func(F) (K, V)) map[K]V {
	ret := make(map[K]V, len(s))
	for _, e := range s {
		k, v := fc(e)
		ret[k] = v
	}
	return ret
}

// Filter returns a new slice containing only the elements of s for which fc returns true.
//
// EXAMPLE:
//
//	Filter([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }) => []int{2, 4}
//	Filter([]int{}, func(i int) bool { return i%2 == 0 })           => []int{}
//	Filter(nil, func(i int) bool { return i%2 == 0 })              => []int{}
//
// HINT:
//
//   - Use [Reject] if you want to exclude elements for which fc returns true.
func Filter[F any](s []F, fc func(F) bool) []F {
	ret := make([]F, 0, len(s)/2)
	for _, v := range s {
		if fc(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// All returns true if fc returns true for all elements of s.
// Returns false if any element fails the condition.
//
// EXAMPLE:
//
//	All([]int{2, 4, 6}, func(i int) bool { return i%2 == 0 }) => true
//	All([]int{2, 3, 6}, func(i int) bool { return i%2 == 0 }) => false
//	All([]int{}, func(i int) bool { return i%2 == 0 })        => true
//
// HINT:
//
//   - Use [Any] if you want to check if at least one element satisfies the condition.
func All[T any](s []T, fc func(T) bool) bool {
	for _, v := range s {
		if !fc(v) {
			return false
		}
	}
	return true
}

// Any returns true if fc returns true for at least one element of s.
// Returns false if no element satisfies the condition.
//
// EXAMPLE:
//
//	Any([]int{2, 4, 6}, func(i int) bool { return i%2 != 0 }) => false
//	Any([]int{2, 3, 6}, func(i int) bool { return i%2 != 0 }) => true
//	Any([]int{}, func(i int) bool { return i%2 != 0 })        => false
//
// HINT:
//
//   - Use [All] if you want to check if all elements satisfy the condition.
func Any[T any](s []T, fc func(T) bool) bool {
	for _, v := range s {
		if fc(v) {
			return true
		}
	}
	return false
}

// First returns the first element in slice s that satisfies the condition fc.
// If no element satisfies the condition, it returns the zero value of T and false.
//
// EXAMPLE:
//
//	First([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 }) => (2, true)
//	First([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 }) => (0, false)
//	First([]int{}, func(i int) bool { return i%2 == 0 })        => (0, false)
//
// HINT:
//
//   - Use [Last] if you want the last element that satisfies the condition.
func First[T any, S ~[]T](s S, fc func(T) bool) (T, bool) {
	for _, v := range s {
		if fc(v) {
			return v, true
		}
	}
	return gvalue.Zero[T](), false
}

// Last returns the last element in slice s that satisfies the condition fc.
// If no element satisfies the condition, it returns the zero value of T and false.
//
// EXAMPLE:
//
//	Last([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 }) => (2, true)
//	Last([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 }) => (0, false)
//	Last([]int{}, func(i int) bool { return i%2 == 0 })        => (0, false)
//
// HINT:
//
//   - Use [First] if you want the first element that satisfies the condition.
func Last[T any, S ~[]T](s S, fc func(T) bool) (T, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if fc(s[i]) {
			return s[i], true
		}
	}
	return gvalue.Zero[T](), false
}

// FilterMap applies function fc to each element of slice s with type F.
// If fc returns (T, true), the result is included in the output slice.
// If fc returns (T, false), the result is ignored.
//
// EXAMPLE:
//
//	FilterMap([]int{1, 2, 3, 4}, func(i int) (string, bool) {
//		if i%2 == 0 {
//			return strconv.Itoa(i), true
//		}
//		return "", false
//	}) => []string{"2", "4"}
//
// HINT:
//
//   - Use [Map] if you want to include all elements in the output.
func FilterMap[F, T any](s []F, fc func(F) (T, bool)) []T {
	ret := make([]T, 0, len(s)/2)
	for _, v := range s {
		if t, ok := fc(v); ok {
			ret = append(ret, t)
		}
	}
	return ret
}

// Reject returns a new slice containing only the elements of s for which fc returns false.
//
// EXAMPLE:
//
//	Reject([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }) => []int{1, 3}
//	Reject([]int{}, func(i int) bool { return i%2 == 0 })           => []int{}
//	Reject(nil, func(i int) bool { return i%2 == 0 })              => []int{}
//
// HINT:
//
//   - Use [Filter] if you want to include elements for which fc returns true.
func Reject[T any](s []T, fc func(T) bool) []T {
	ret := make([]T, 0, len(s)/2)
	for _, v := range s {
		if !fc(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Reduce reduces the slice s to a single value by applying function fc to each element.
// The first element of s is used as the initial value.
//
// EXAMPLE:
//
//	Reduce([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }) => 10
//	Reduce([]int{}, func(a, b int) int { return a + b })           => 0
//
// HINT:
//
//   - Use [Fold] if you want to specify an initial value.
func Reduce[T any](s []T, fc func(T, T) T) T {
	if len(s) == 0 {
		return gvalue.Zero[T]()
	}
	ret := s[0]
	for _, v := range s[1:] {
		ret = fc(ret, v)
	}
	return ret
}

// Fold reduces the slice s to a single value by applying function fc to each element.
// The initial value is provided as init.
//
// EXAMPLE:
//
//	Fold([]int{1, 2, 3, 4}, func(a, b int) int { return a + b }, 10) => 20
//	Fold([]int{}, func(a, b int) int { return a + b }, 10)           => 10
//
// HINT:
//
//   - Use [Reduce] if you want to use the first element as the initial value.
func Fold[T1, T2 any](s []T1, fc func(T2, T1) T2, init T2) T2 {
	ret := init
	for _, v := range s {
		ret = fc(ret, v)
	}
	return ret
}

// ForEach applies function fc to each element of slice s.
//
// EXAMPLE:
//
//	ForEach([]int{1, 2, 3}, func(i int) { fmt.Println(i) }) => prints 1, 2, 3
//	ForEach([]int{}, func(i int) { fmt.Println(i) })        => prints nothing
//
// HINT:
//
//   - Use [ForEachIdx] if you also need the index of each element.
func ForEach[T any](s []T, fc func(T)) {
	for _, v := range s {
		fc(v)
	}
}

// ForEachIdx applies function fc to each element of slice s, passing the index and the element.
//
// EXAMPLE:
//
//	ForEachIdx([]int{1, 2, 3}, func(i, v int) { fmt.Printf("%d: %d\n", i, v) }) => prints "0: 1", "1: 2", "2: 3"
//	ForEachIdx([]int{}, func(i, v int) { fmt.Printf("%d: %d\n", i, v) })        => prints nothing
//
// HINT:
//
//   - Use [ForEach] if you don't need the index.
func ForEachIdx[T any](s []T, fc func(int, T)) {
	for i, v := range s {
		fc(i, v)
	}
}

// Find returns the first element in slice s that satisfies the condition f.
// If no element satisfies the condition, it returns the zero value of T and false.
//
// EXAMPLE:
//
//	Find([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 }) => (2, true)
//	Find([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 }) => (0, false)
//	Find([]int{}, func(i int) bool { return i%2 == 0 })        => (0, false)
//
// HINT:
//
//   - Use [FindRev] if you want the last element that satisfies the condition.
func Find[T any](s []T, f func(T) bool) (T, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	return gvalue.Zero[T](), false
}

// FindRev returns the last element in slice s that satisfies the condition f.
// If no element satisfies the condition, it returns the zero value of T and false.
//
// EXAMPLE:
//
//	FindRev([]int{1, 2, 3}, func(i int) bool { return i%2 == 0 }) => (2, true)
//	FindRev([]int{1, 3, 5}, func(i int) bool { return i%2 == 0 }) => (0, false)
//	FindRev([]int{}, func(i int) bool { return i%2 == 0 })        => (0, false)
//
// HINT:
//
//   - Use [Find] if you want the first element that satisfies the condition.
func FindRev[T any](s []T, f func(T) bool) (T, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i], true
		}
	}
	return gvalue.Zero[T](), false
}

// GroupBy groups elements of slice s by the key returned by function f.
// Returns a map where the keys are the results of f and the values are slices of elements.
//
// EXAMPLE:
//
//	GroupBy([]int{1, 2, 3, 4}, func(i int) int { return i % 2 }) => map[int][]int{0: {2, 4}, 1: {1, 3}}
//	GroupBy([]int{}, func(i int) int { return i % 2 })          => map[int][]int{}
//
// HINT:
//
//   - Ensure that the key type K is comparable.
func GroupBy[K comparable, T any, S ~[]T](s S, f func(T) K) map[K]S {
	m := make(map[K]S)
	for i := range s {
		k := f(s[i])
		m[k] = append(m[k], s[i])
	}
	return m
}

// Contains returns true if slice s contains the value v.
//
// EXAMPLE:
//
//	Contains([]int{1, 2, 3}, 2) => true
//	Contains([]int{1, 2, 3}, 4) => false
//	Contains([]int{}, 2)        => false
//
// HINT:
//
//   - Use [ContainsAll] or [ContainsAny] for multiple values.
func Contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if v == vv {
			return true
		}
	}
	return false
}

// ContainsAll returns true if slice s contains all values in vs.
//
// EXAMPLE:
//
//	ContainsAll([]int{1, 2, 3}, 2, 3) => true
//	ContainsAll([]int{1, 2, 3}, 2, 4) => false
//	ContainsAll([]int{}, 2, 3)        => false
//
// HINT:
//
//   - Use [ContainsAny] if you want to check for at least one value.
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

// ContainsAny returns true if slice s contains at least one value in vs.
//
// EXAMPLE:
//
//	ContainsAny([]int{1, 2, 3}, 2, 4) => true
//	ContainsAny([]int{1, 2, 3}, 4, 5) => false
//	ContainsAny([]int{}, 2, 3)        => false
//
// HINT:
//
//   - Use [ContainsAll] if you want to check for all values.
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

// Remove returns a new slice with all occurrences of value v removed from slice s.
//
// EXAMPLE:
//
//	Remove([]int{1, 2, 3, 2}, 2) => []int{1, 3}
//	Remove([]int{1, 2, 3}, 4)    => []int{1, 2, 3}
//	Remove([]int{}, 2)           => []int{}
//
// HINT:
//
//   - Use [RemoveN] if you want to remove only a specific number of occurrences.
func Remove[T comparable](s []T, v T) []T {
	return Filter(s, func(t T) bool {
		return t != v
	})
}

// RemoveN returns a new slice with up to n occurrences of value v removed from slice s.
//
// EXAMPLE:
//
//	RemoveN([]int{1, 2, 3, 2, 2}, 2, 2) => []int{1, 3, 2}
//	RemoveN([]int{1, 2, 3}, 4, 1)       => []int{1, 2, 3}
//	RemoveN([]int{}, 2, 1)              => []int{}
//
// HINT:
//
//   - Use [Remove] if you want to remove all occurrences of v.
func RemoveN[T comparable](s []T, v T, n int) []T {
	return Filter(s, func(t T) bool {
		if n <= 0 {
			return true
		}
		if t == v {
			n--
			return false
		}
		return true
	})
}

// Distinct returns a new slice with duplicate elements removed from slice s.
//
// EXAMPLE:
//
//	Distinct([]int{1, 2, 2, 3, 3, 3}) => []int{1, 2, 3}
//	Distinct([]int{})                => []int{}
//	Distinct(nil)                    => []int{}
//
// HINT:
//
//   - Use [DistinctBy] if you want to remove duplicates based on a custom key.
func Distinct[T comparable](s []T) []T {
	return gutils.NewSliceSetFormSlice(s).GetSlice()
}

// DistinctBy returns a new slice with duplicate elements removed based on the key returned by function fc.
//
// EXAMPLE:
//
//	DistinctBy([]string{"apple", "banana", "apricot"}, func(s string) string {
//		return string(s[0])
//	}) => []string{"apple", "banana"}
//
// HINT:
//
//   - Use [Distinct] if you want to remove duplicates based on the element itself.
func DistinctBy[K comparable, V any](s []V, fc func(V) K) []V {
	ss := gutils.NewSliceSet[K, V]()
	for _, v := range s {
		ss.Upsert(fc(v), v)
	}
	return ss.GetSlice()
}

// DeepCopy returns a deep copy of slice s.
//
// EXAMPLE:
//
//	DeepCopy([]int{1, 2, 3}) => []int{1, 2, 3}
//	DeepCopy([]int{})        => []int{}
//	DeepCopy(nil)            => nil
//
// HINT:
//
//   - Use this function to ensure that modifications to the returned slice do not affect the original slice.
func DeepCopy[T any, S ~[]T](s S) S {
	if s == nil {
		return nil
	}
	return Map(s, func(t T) T { return t })
}

// LastE returns the last element of slice s.
// If the slice is empty, it returns the zero value of T.
//
// EXAMPLE:
//
//	LastE([]int{1, 2, 3}) => 3
//	LastE([]int{})        => 0
//	LastE(nil)            => 0
//
// HINT:
//
//   - Use [Last] if you want to find the last element that satisfies a condition.
func LastE[T any](s []T) T {
	if len(s) == 0 {
		return gvalue.Zero[T]()
	}
	return s[len(s)-1]
}

// Union returns a new slice containing the union of all input slices, with duplicates removed.
//
// EXAMPLE:
//
//	Union([]int{1, 2}, []int{2, 3}, []int{3, 4}) => []int{1, 2, 3, 4}
//	Union([]int{})                               => []int{}
//	Union(nil)                                   => []int{}
//
// HINT:
//
//   - Use this function to combine multiple slices into one without duplicates.
func Union[K comparable](ss ...[]K) []K {
	if len(ss) == 0 {
		return []K{}
	}
	return gutils.NewSliceSetFormSlice(ss...).GetSlice()
}

// MinMaxBy returns the minimum and maximum elements of slice s based on the comparison function less.
// If the slice is empty, it returns the zero value of T for both minimum and maximum.
//
// EXAMPLE:
//
//	MinMaxBy([]int{3, 1, 4, 1, 5}, func(a, b int) bool { return a < b }) => (1, 5)
//	MinMaxBy([]int{}, func(a, b int) bool { return a < b })              => (0, 0)
//
// HINT:
//
//   - Ensure that the comparison function less is consistent and transitive.
func MinMaxBy[T any](s []T, less func(T, T) bool) (T, T) {
	if len(s) == 0 {
		return gvalue.Zero[T](), gvalue.Zero[T]()
	}
	mins, maxs := s[0], s[0]
	for _, v := range s[1:] {
		if less(v, mins) {
			mins = v
		} else if less(maxs, v) {
			maxs = v
		}
	}
	return mins, maxs
}

// UnsafeAsserts converts a slice of type From to a slice of type To using unsafe type assertions.
//
// EXAMPLE:
//
//	UnsafeAsserts[float64, int]([]int{1, 2, 3}) => []float64{1.0, 2.0, 3.0}
//
// HINT:
//
//   - Use [SafeAsserts] if you want to avoid panics due to invalid type assertions.
func UnsafeAsserts[To, From any](s []From) []To {
	return Map[From, To](s, func(from From) To {
		return any(from).(To)
	})
}

// SafeAsserts converts a slice of type From to a slice of type To using safe type assertions.
// If a type assertion fails, the zero value of To is used.
//
// EXAMPLE:
//
//	SafeAsserts[float64, any]([]any{1, "2", 3.0}) => []float64{1.0, 0.0, 3.0}
//
// HINT:
//
//   - Use [UnsafeAsserts] if you are certain that all type assertions will succeed.
func SafeAsserts[To, From any](s []From) []To {
	return Map[From, To](s, func(from From) To {
		return gvalue.SafeAssert[To](from)
	})
}

// Equal returns true if slices s1 and s2 are equal (contain the same elements in the same order).
//
// EXAMPLE:
//
//	Equal([]int{1, 2, 3}, []int{1, 2, 3}) => true
//	Equal([]int{1, 2, 3}, []int{3, 2, 1}) => false
//	Equal([]int{}, []int{})               => true
//
// HINT:
//
//   - Use [EqualStrict] if you want to also compare nil slices.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualBy returns true if slices s1 and s2 are equal based on the comparison function eq.
//
// EXAMPLE:
//
//	EqualBy([]int{1, 2, 3}, []int{2, 3, 4}, func(a, b int) bool { return a+1 == b }) => true
//	EqualBy([]int{1, 2, 3}, []int{3, 2, 1}, func(a, b int) bool { return a == b })   => false
//	EqualBy([]int{}, []int{}, func(a, b int) bool { return a == b })                => true
//
// HINT:
//
//   - Use [Equal] if you want to compare elements directly.
func EqualBy[T any](s1, s2 []T, eq func(T, T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

// EqualStrict returns true if slices s1 and s2 are strictly equal, including nil checks.
//
// EXAMPLE:
//
//	EqualStrict([]int{1, 2, 3}, []int{1, 2, 3}) => true
//	EqualStrict(nil, nil)                       => true
//	EqualStrict([]int{}, nil)                   => false
//
// HINT:
//
//   - Use [Equal] if you don't need to distinguish between nil and empty slices.
func EqualStrict[T comparable](s1, s2 []T) bool {
	if (s1 == nil && s2 != nil) || (s1 != nil && s2 == nil) {
		return false
	}
	return Equal(s1, s2)
}
