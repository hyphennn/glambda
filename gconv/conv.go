// Package gconv
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/20
package gconv

import (
	"github.com/hyphennn/glambda/internal"
	"github.com/hyphennn/glambda/internal/constraints"
)

func ToPtr[T constraints.UnPtrAble](t T) *T {
	return &t
}

func FromPtr[T constraints.UnPtrAble](t *T) T {
	if t == nil {
		return internal.Zero[T]()
	}
	return *t
}

func StringPtr(s string) *string {
	return &s
}

func Ptr2String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
