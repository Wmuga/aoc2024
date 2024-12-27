package set

import "github.com/wmuga/aoc2019/pkg/set"

type Set[T comparable] struct {
	*set.Set[T]
}

func (s *Set[T]) Clone() *Set[T] {
	newSet := &Set[T]{Set: set.NewSet[T]()}

	for v := range s.Set.Iterator() {
		newSet.Set.Upsert(v)
	}

	return newSet
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.Set.Get(item)
	return ok
}

func New[T comparable]() *Set[T] {
	return &Set[T]{Set: set.NewSet[T]()}
}
