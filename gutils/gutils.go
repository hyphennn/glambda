// Package gutils
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gutils

func TernaryForm[T any](cond bool, tureVal, falseVal T) T {
	if cond {
		return tureVal
	}
	return falseVal
}
