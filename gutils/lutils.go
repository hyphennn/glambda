// Package gutils
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gutils

import (
	"context"
	"sync"

	"github.com/hyphennn/glamda/internal"
)

func TernaryForm[T any](cond bool, tureVal, falseVal T) T {
	if cond {
		return tureVal
	}
	return falseVal
}

type Pair[F, S any] struct {
	First  F
	Second S
}

func MakePair[F, S any](f F, s S) *Pair[F, S] {
	return &Pair[F, S]{First: f, Second: s}
}

func (p *Pair[F, S]) Split() (F, S) {
	return p.First, p.Second
}

func FastAssert[T any](v any) T {
	t, ok := v.(T)
	if !ok {
		return internal.Zero[T]()
	}
	return t
}

func MustDo[K, V any](key K, fc func(K) (V, error)) V {
	return MustEasyDo(func() (V, error) {
		return fc(key)
	})
}

func MustDoCtx[K, V any](ctx context.Context, key K, fc func(context.Context, K) (V, error)) V {
	return MustEasyDo(func() (V, error) {
		return fc(ctx, key)
	})
}

func MustEasyDo[V any](fc func() (V, error)) V {
	v, err := fc()
	if err != nil {
		return internal.Zero[V]()
	}
	return v
}

type ch[T any] chan T

type SafeChan[T any] struct {
	ch[T]
	once sync.Once
}

func NewSafeChan[T any](size ...int) *SafeChan[T] {
	if len(size) == 0 {
		return &SafeChan[T]{make(chan T), sync.Once{}}
	}
	return &SafeChan[T]{make(chan T, size[0]), sync.Once{}}
}

func (s *SafeChan[T]) Listen() (t T) {
	t = <-s.ch
	return
}

func (s *SafeChan[T]) Send(t T) {
	s.ch <- t
}

func (s *SafeChan[T]) Close() {
	s.once.Do(func() {
		close(s.ch)
	})
}

func Paging[T any](arr []T, offset, limit int) []T {
	return arr[TernaryForm((offset)*limit <= len(arr), (offset)*limit, len(arr)):TernaryForm((offset+1)*limit <= len(arr), (offset+1)*limit, len(arr))]
}
