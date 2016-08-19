//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package intset

import (
	"fmt"
	"strings"
	"sort"
)

type MapIntSet map[int]bool

func (s *MapIntSet) Has(x int) bool {
	return (*s)[x]
}

func (s *MapIntSet) Add(x int) {
	(*s)[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for k := range *t {
		(*s)[k] = true
	}
}

func (s *MapIntSet) IntersectWith(t *MapIntSet) {
	iw := make(MapIntSet)
	for k := range *t {
		if (*s)[k] {
			iw[k] = true
		}
	}
	*s = iw
}

func (s *MapIntSet) DifferenceWith(t *MapIntSet) {
	dw := make(MapIntSet)
	for k := range *s {
		if !(*t)[k] {
			dw[k] = true
		}
	}
	*s = dw
}

func (s *MapIntSet) Elems() (list []int) {
	for k := range *s {
		list = append(list, k)
	}
	return list
}

// (1,1)->0, (1,0)->1 (0,1)->1 (0,0)->0
func (s *MapIntSet) SymmetricDifference(t *MapIntSet) {
	for k := range *t {
		if (*s)[k] {
			(*s)[k] = true
		} else {
			delete(*s, k)
		}
	}
}

func (s *MapIntSet) String() string {
	is := []int{}
	for k := range *s {
		is = append(is, k)
	}
	sort.Sort(sort.IntSlice(is))

	str := "{"
	for _, i := range is {
		str += fmt.Sprintf("%d ", i)
	}
	str = strings.TrimSuffix(str, " ")
	return str + "}"
}

func (s *MapIntSet) Len() int {
	return len(*s)
}

func (s *MapIntSet) Remove(x int) {
	delete(*s, x)
}

func (s *MapIntSet) Clear() {
	*s = make(MapIntSet)
}

func (s *MapIntSet) Copy() *MapIntSet {
	clone := make(MapIntSet)
	for k, v := range *s {
		clone[k] = v
	}
	return &clone
}

func (s *MapIntSet) AddAll(x ...int) {
	for _, xx := range x {
		s.Add(xx)
	}
}
