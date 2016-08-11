//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package intset

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var setSize int = 200

func BenchmarkAdd(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	set := IntSet{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < setSize; j++ {
			set.Add(rng.Intn(math.MaxInt32))
		}
		set.Clear()
	}
}

func BenchmarkMapAdd(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	set := MapIntSet{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < setSize; j++ {
			set.Add(rng.Intn(math.MaxInt32))
		}
		set.Clear()
	}
}

func TestMapIntSetAdd(t *testing.T) {
	testData := [][]int{
		[]int{0, bitSize, bitSize / 2},
		[]int{0, 2, 100},
	}
	for _, td := range testData {
		set := MapIntSet{}
		for _, x := range td {
			set.Add(x)
		}
		if set.Len() != len(td) {
			t.Errorf("MapIntSet Len is unmatch. actual = %d. expected = %d.\n", set.Len(), len(td))
		}
		for _, x := range td {
			if !set.Has(x) {
				t.Errorf("MapIntSet don't have %q. data = %q\n", x, td)
			}
		}
	}
}

func TestIntSetAdd(t *testing.T) {
	testData := [][]int{
		[]int{0, bitSize, bitSize / 2},
		[]int{0, 2, 100},
	}
	for _, td := range testData {
		set := IntSet{}
		for _, x := range td {
			set.Add(x)
		}
		if set.Len() != len(td) {
			t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d.\n", set.Len(), len(td))
		}
		for _, x := range td {
			if !set.Has(x) {
				t.Errorf("IntSet don't have %q. data = %q\n", x, td)
			}
		}
	}
}

func TestIntSetString(t *testing.T) {
	testData := []struct {
		data     []int
		expected string
	}{
		{
			[]int{0, 2, 100},
			"{0 2 100}",
		},
	}
	for _, td := range testData {
		set := IntSet{}
		for _, x := range td.data {
			set.Add(x)
		}
		if set.String() != td.expected {
			t.Errorf("actual = %s \nexpected = %s\n", set.String(), td.expected)
		}
	}
}

func TestMapIntSetString(t *testing.T) {
	testData := []struct {
		data     []int
		expected string
	}{
		{
			[]int{0, 2, 100},
			"{0 2 100}",
		},
	}
	for _, td := range testData {
		set := MapIntSet{}
		for _, x := range td.data {
			set.Add(x)
		}
		if set.String() != td.expected {
			t.Errorf("actual = %s \nexpected = %s\n", set.String(), td.expected)
		}
	}
}

func TestIntSetUnionWith(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{0, 1, bitSize, bitSize + 1, bitSize / 2}
	data1 := IntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := IntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.UnionWith(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d.\n", data1.Len(), len(expected))
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %q. data = %q\n", x, expected)
		}
	}
}

func TestMapIntSetUnionWith(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{0, 1, bitSize, bitSize + 1, bitSize / 2}
	data1 := MapIntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := MapIntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.UnionWith(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d.\n", data1.Len(), len(expected))
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %q. data = %q\n", x, expected)
		}
	}
}

func TestMapIntIntersectWith(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{bitSize / 2}
	data1 := MapIntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := MapIntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.IntersectWith(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d. data1 = %q\n", data1.Len(), len(expected), data1.String())
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %q. data = %q\n", x, expected)
		}
	}
}

func TestIntIntersectWith(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{bitSize / 2}
	data1 := IntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := IntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.IntersectWith(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d. data1 = %q\n", data1.Len(), len(expected), data1.String())
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %q. data = %q\n", x, expected)
		}
	}
}

func TestIntDifferentWith(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{0, bitSize}
	data1 := IntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := IntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.DifferenceWith(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d. data1 = %q\n", data1.Len(), len(expected), data1.String())
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %q. data = %q\n", x, data1.String())
		}
	}
}

func TestMapIntDifferentWith(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{0, bitSize}
	data1 := MapIntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := MapIntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.DifferenceWith(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d. data1 = %q\n", data1.Len(), len(expected), data1.String())
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %d. data = %q\n", x, data1.String())
		}
	}
}

/*
func TestSymmetricDifferenfce(t *testing.T) {
	testData1 := []int{0, bitSize, bitSize / 2}
	testData2 := []int{1, bitSize + 1, bitSize / 2}
	expected := []int{0, 1, bitSize, bitSize + 1}
	data1 := IntSet{}
	for _, td := range testData1 {
		data1.Add(td)
	}
	data2 := IntSet{}
	for _, td := range testData2 {
		data2.Add(td)
	}
	data1.SymmetricDifference(&data2)
	if data1.Len() != len(expected) {
		t.Errorf("IntSet Len is unmatch. actual = %d. expected = %d. data1 = %q\n", data1.Len(), len(expected), data1.String())
	}
	for _, x := range expected {
		if !data1.Has(x) {
			t.Errorf("IntSet don't have %d. data = %q\n", x, data1.String())
		}
	}
}
*/

func TestElems(t *testing.T) {
	testData1 := []int{0, bitSize / 2, bitSize}
	data := IntSet{}
	for _, td := range testData1 {
		data.Add(td)
	}
	actual := data.Elems()
	if len(actual) != len(testData1) {
		t.Errorf("IntSet len(IntSet.Elems()) is unmatch. actual = %d. expected = %d. data1 = %q\n", len(actual), len(testData1), data.String())
	}

	sort.Ints(testData1)
	sort.Ints(actual)
	for i := 0; i < len(testData1); i++ {
		if testData1[i] != actual[i] {
			t.Errorf("actual = %q, expected = %q", actual, testData1)
		}
	}
}

func TestMapElems(t *testing.T) {
	testData1 := []int{0, bitSize / 2, bitSize}
	data := MapIntSet{}
	for _, td := range testData1 {
		data.Add(td)
	}
	actual := data.Elems()
	if len(actual) != len(testData1) {
		t.Errorf("IntSet len(IntSet.Elems()) is unmatch. actual = %d. expected = %d. data1 = %q\n", len(actual), len(testData1), data.String())
	}

	sort.Ints(testData1)
	sort.Ints(actual)
	for i := 0; i < len(testData1); i++ {
		if testData1[i] != actual[i] {
			t.Errorf("actual = %q, expected = %q", actual, testData1)
		}
	}
}

func TestMapRemove(t *testing.T) {
	testData := []int{0, bitSize / 2, bitSize}
	removeData := []int{bitSize / 2}
	expected := []int{0, bitSize}
	data := MapIntSet{}
	for _, td := range testData {
		data.Add(td)
	}
	for _, rd := range removeData {
		data.Remove(rd)
	}

	if data.Len() != len(expected) {
		t.Errorf("IntSet len is unmatch. actual = %d. expected = %d. data1 = %q\n", data.Len(), len(expected), data.String())
	}

	for _, x := range expected {
		if !data.Has(x) {
			t.Errorf("IntSet don't have %d. data = %q\n", x, data.String())
		}
	}
}

func TestRemove(t *testing.T) {
	testData := []int{0, bitSize / 2, bitSize}
	removeData := []int{bitSize / 2}
	expected := []int{0, bitSize}
	data := IntSet{}
	for _, td := range testData {
		data.Add(td)
	}
	for _, rd := range removeData {
		data.Remove(rd)
	}

	if data.Len() != len(expected) {
		t.Errorf("IntSet len is unmatch. actual = %d. expected = %d. data1 = %q\n", data.Len(), len(expected), data.String())
	}

	for _, x := range expected {
		if !data.Has(x) {
			t.Errorf("IntSet don't have %d. data = %q\n", x, data.String())
		}
	}
}

func TestCopy(t *testing.T) {
	testData := []int{0, bitSize / 2, bitSize}

	data := IntSet{}
	for _, td := range testData {
		data.Add(td)
	}
	copied := data.Copy()

	if &data == copied {
		t.Errorf("Address is coppied.\n")
	}

	comp1 := copied.Elems()
	comp2 := testData
	sort.Ints(comp1)
	sort.Ints(comp2)
	for i := 0; i < len(comp1); i++ {
		if comp1[i] != comp2[i] {
			t.Errorf("actual = %q, expected = %q", comp1, comp2)
		}
	}
}

func TestMapCopy(t *testing.T) {
	testData := []int{0, bitSize / 2, bitSize}

	data := MapIntSet{}
	for _, td := range testData {
		data.Add(td)
	}
	copied := data.Copy()

	if &data == copied {
		t.Errorf("Address is coppied.\n")
	}

	comp1 := copied.Elems()
	comp2 := testData
	sort.Ints(comp1)
	sort.Ints(comp2)
	for i := 0; i < len(comp1); i++ {
		if comp1[i] != comp2[i] {
			t.Errorf("actual = %q, expected = %q", comp1, comp2)
		}
	}
}

func TestMapAddAll(t *testing.T) {
	testData := []int{0, bitSize / 2, bitSize}

	data := MapIntSet{}
	data.AddAll(testData...)
	for _, x := range testData {
		if !data.Has(x) {
			t.Errorf("IntSet don't have %d. data = %q\n", x, data.String())
		}
	}
}

func TestAddAll(t *testing.T) {
	testData := []int{0, bitSize / 2, bitSize}

	data := IntSet{}
	data.AddAll(testData...)
	for _, x := range testData {
		if !data.Has(x) {
			t.Errorf("IntSet don't have %d. data = %q\n", x, data.String())
		}
	}
}
