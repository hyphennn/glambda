// Package disjoint_set
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/9/30
package disjoint_set

type DisjointSet[T comparable] interface {
	Find(value T) (T, error)
	IsExist(value T) bool
	Merge(a, b T) (bool, error)
	AppendNode(value T)
}
