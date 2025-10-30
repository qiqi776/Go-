package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// set基于map定义了一个存放元素的集合类型
type Set[E comparable] struct {
	m map[E]struct{}
}

func NewSet[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}

func (s *Set[E]) Add(e E) {
	s.m[e] = struct{}{}
}

// Push迭代器
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

func forRangeSet() {
	s := NewSet[string]()
	s.Add("Golang")
	s.Add("Java")
	s.Add("Python")
	s.Add("C++")
	for v := range s.All() {
		fmt.Println(v)
	}
}

// Pull迭代器
func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, ok1 := next()
			if !ok1 {
				return
			}
			v2, ok2 := next()
			if !yield(v1, v2) {
				return
			}
			if !ok2 {
				return
			}
		}
	}
}

func SortDemo() {
	m := map[int]string{
		1: "Golang",
		2: "Java",
		3: "Python",
	}

	for _, key := range slices.Sorted(maps.Keys(m)) {
		fmt.Println(m[key])
	}
}
