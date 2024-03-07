// Package gconv
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/20
package gconv_test

import (
	"testing"

	"github.com/hyphennn/glambda/gconv"
	"github.com/hyphennn/glambda/internal/assert"
)

func TestToPtr(t *testing.T) {
	s := gconv.ToPtr("123")
	assert.Equal(t, "123", *s)
}

func TestFromPtr(t *testing.T) {
	s := "123"
	assert.Equal(t, "123", gconv.FromPtr(&s))
}
