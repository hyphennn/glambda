// Package gutils
// Author: hyphen
// Copyright 2024 hyphen. All rights reserved.
// Create-time: 2024/2/26
package gutils

import (
	"testing"
)

func TestURLFmt(t *testing.T) {
	type args struct {
		pattern string
		args    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"normal case",
			args{
				pattern: "/test/{a}/{b}/ccc/{d}",
				args:    []string{"{a}", "aaa", "{b}", "bbb", "{d}", "ddd"},
			},
			"/test/aaa/bbb/ccc/ddd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLFmt(tt.args.pattern, tt.args.args...); got != tt.want {
				t.Errorf("URLFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
