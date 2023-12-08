// Package disjoint_set
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/10/2
package disjoint_set

type DisjointSetWithoutPC[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}
