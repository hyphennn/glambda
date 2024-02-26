// Package gutils
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/12/4
package gutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

var client = &http.Client{}

func AccessResp(ctx context.Context, url string, method string, header map[string]string, param map[string]string,
	body any, setAuthorization func(*http.Request)) (*http.Response, error) {
	return access(ctx, url, method, header, param, body, setAuthorization)
}

func Access[T any](ctx context.Context, url string, method string, header map[string]string, param map[string]string,
	body any, setAuthorization func(*http.Request), isSuccess func(response *http.Response) bool) (T, error) {
	var ret T
	resp, err := access(ctx, url, method, header, param, body, setAuthorization)
	if err != nil {
		return ret, fmt.Errorf("access %s failed: %w", url, err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, fmt.Errorf("read resp body failed: %w", err)
	}
	if isSuccess != nil && !isSuccess(resp) {
		return ret, fmt.Errorf("is success return false: %s", string(respBody))
	}
	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		return ret, fmt.Errorf("unmarshal resp body failed: %w", err)
	}
	return ret, nil
}

func access(ctx context.Context, url string, method string, header map[string]string, param map[string]string,
	body any, setAuthorization func(*http.Request)) (*http.Response, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyJSON)
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if setAuthorization != nil {
		setAuthorization(req)
	}
	q := req.URL.Query()
	for k, v := range param {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("access %s failed: %w", url, err)
	}
	return resp, nil
}

func URLFmt(pattern string, args ...string) string {
	if len(args)%2 == 1 {
		return pattern
	}
	for i := 0; i < len(args); i += 2 {
		pattern = strings.ReplaceAll(pattern, args[i], args[i+1])
	}
	return pattern
}
