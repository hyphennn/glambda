// Package lset
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/8
package lset

type SliceSet[K comparable, V any] struct {
	m     map[K]int
	slice []V
}

func NewSliceSet[K comparable, V any]() *SliceSet[K, V] {
	return &SliceSet[K, V]{map[K]int{}, []V{}}
}

func NewSliceSetFormSlice[K comparable](s []K) *SliceSet[K, K] {
	ret := &SliceSet[K, K]{map[K]int{}, []K{}}
	for _, k := range s {
		ret.Upsert(k, k)
	}
	return ret
}

func (s *SliceSet[K, V]) insert(key K, value V) {
	s.m[key] = len(s.slice)
	s.slice = append(s.slice, value)
}

func (s *SliceSet[K, V]) update(key K, value V) {
	s.slice[s.m[key]] = value
}

func (s *SliceSet[K, V]) Insert(key K, value V) bool {
	if _, ok := s.m[key]; ok {
		return false
	}
	s.insert(key, value)
	return true
}

func (s *SliceSet[K, V]) Update(key K, value V) bool {
	if _, ok := s.m[key]; !ok {
		return false
	}
	s.update(key, value)
	return true
}
func (s *SliceSet[K, V]) Upsert(key K, value V) {
	if _, ok := s.m[key]; ok {
		s.update(key, value)
		return
	}
	s.insert(key, value)
}

func (s *SliceSet[K, V]) Get(key K) (V, bool) {
	var v V
	i, ok := s.m[key]
	if !ok {
		return v, ok
	}
	return s.slice[i], ok
}

func (s *SliceSet[K, V]) GetSlice() []V {
	return s.slice
}

func (s *SliceSet[K, V]) GetMap() map[K]int {
	return s.m
}
