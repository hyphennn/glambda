// Package disjoint_set
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/8/5
package disjoint_set

// todo 能否对启发式优化进行自定义设置？
import (
	"fmt"

	"github.com/hyphennn/glambda/gutils"
)

type DisjointSetWithPC[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

func NewDisjointSetWithPC[T comparable](values []T, enablePC ...bool) *DisjointSetWithPC[T] {
	parent, rank := map[T]T{}, map[T]int{}
	for _, value := range values {
		parent[value] = value
		rank[value] = 0
	}
	return &DisjointSetWithPC[T]{parent, rank}
}

func (d *DisjointSetWithPC[T]) Find(value T) (T, error) {
	var v T
	if !d.IsExist(value) {
		return v, fmt.Errorf("[*DisjointSetWithPC.Find]value doesn't exist in disjoint set")
	}
	return d.innerFind(value), nil
}

func (d *DisjointSetWithPC[T]) innerFind(value T) T {
	return gutils.TernaryForm(d.parent[value] != value, func() T {
		d.parent[value] = d.innerFind(d.parent[value])
		return value
	}(), d.parent[value])
}

func (d *DisjointSetWithPC[T]) IsExist(value T) bool {
	_, ok := d.parent[value]
	return ok
}

func (d *DisjointSetWithPC[T]) Merge(a, b T) (bool, error) {
	if !d.IsExist(a) {
		return false, fmt.Errorf("[*DisjointSetWithPC.Merge]a doesn't exist in disjoint set")
	}
	if !d.IsExist(b) {
		return false, fmt.Errorf("[*DisjointSetWithPC.Merge]b doesn't exist in disjoint set")
	}
	aRoot, bRoot := d.innerFind(a), d.innerFind(b)
	if aRoot == bRoot {
		return false, nil
	}
	if d.rank[aRoot] > d.rank[bRoot] {
		d.parent[bRoot] = aRoot
	} else if d.rank[aRoot] < d.rank[bRoot] {
		d.parent[aRoot] = bRoot
	} else {
		d.rank[aRoot]++
		d.parent[bRoot] = aRoot
	}
	return true, nil
}

func (d *DisjointSetWithPC[T]) AppendNode(value T) error {
	if d.IsExist(value) {
		return fmt.Errorf("[*DisjointSetWithPC.AppendNode]value exists in disjoint set")
	}
	d.parent[value] = value
	return nil
}
