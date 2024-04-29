package main

import (
	process "github.com/ogios/ansisgr-process"
	"github.com/ogios/cropviewport"
)

type ContentCacher[T any] struct {
	pool map[string]T
	new  func(p string) T
}

func (c *ContentCacher[T]) Get(key string) T {
	if s, ok := c.pool[key]; ok {
		return s
	}
	s := c.new(key)
	c.pool[key] = s
	return s
}

func NewStringCacher[T any](new func(p string) T) *ContentCacher[T] {
	return &ContentCacher[T]{pool: map[string]T{}, new: new}
}

type ContentData struct {
	Raw   string
	Table *process.ANSITableList
	Lines []*cropviewport.SubLine
}
