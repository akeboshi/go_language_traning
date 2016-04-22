package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntSetAdd(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	checkUint(1, data.words[0], "IntSet.Add(0)", t)
	data.Add(bitSize)
	checkUint(1, data.words[1], "IntSet.Add(uintSize)", t)
	data.Add(bitSize / 2)
	checkUint(1<<uint(bitSize/2)+1, data.words[0], "IntSet.Add(uintSize/2)", t)
}

func TestIntSetString(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	s := fmt.Sprintf("{%d %d %d}", 0, bitSize/2, bitSize)
	checkString(s, data.String(), "IntSet{0,bitSize/2,bitSize}.String", t)
}

func TetIntSetHas(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	checkTrue(data.Has(0), "IntSet.Has(0)", t)
	checkTrue(data.Has(bitSize), "IntSet.Has(bitSize)", t)
	checkTrue(data.Has(bitSize/2), "IntSet.Has(bitSize/2)", t)
	checkFalse(data.Has(1), "IntSet.Has(1)", t)
	checkFalse(data.Has(1000), "IntSet.Has(1000)", t)
}

func TestIntSetLen(t *testing.T) {
	data := IntSet{}
	checkInt(0, data.Len(), "IntSet{NO DATA}.Len()", t)
	data.Add(0)
	data.Add(bitSize / 2)
	data.Add(bitSize)
	checkInt(3, data.Len(), "IntSet{0,bitSize/2,bitSize}.Len()", t)
}

func TestIntSetUnionWith(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	data2 := IntSet{}
	data2.Add(1)
	data2.Add(bitSize + 1)
	data2.Add(bitSize / 2)
	data.UnionWith(&data2)
	checkInt(data.Len(), 5, "UnionWith{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}", t)
	checkTrues("Has{0,1,bitSize,bitSize+1,bitSize/2}", t,
		data.Has(0),
		data.Has(1),
		data.Has(bitSize),
		data.Has(bitSize+1),
		data.Has(bitSize/2))
}

func TestIntIntersectWith(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	data2 := IntSet{}
	data2.Add(1)
	data2.Add(bitSize + 1)
	data2.Add(bitSize / 2)
	data.IntersectWith(&data2)
	checkInt(1, data.Len(), "IntersectWith{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}.Len()", t)
	checkTrues("IntersectWith{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}.Has(biSize/2)", t,
		data.Has(bitSize/2))
}

func TestDifferenceWith(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	data2 := IntSet{}
	data2.Add(1)
	data2.Add(bitSize + 1)
	data2.Add(bitSize / 2)
	data.DifferenceWith(&data2)
	checkInt(2, data.Len(), "DifferenceWith{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}.Len()", t)
	checkTrues("DifferenceWith{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}.Has()", t,
		data.Has(0),
		data.Has(bitSize))
}

func TestSymmetricDifferenfce(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	data2 := IntSet{}
	data2.Add(1)
	data2.Add(bitSize + 1)
	data2.Add(bitSize / 2)
	data.SymmetricDifference(&data2)
	checkInt(4, data.Len(), "SymmetricDifferenfce{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}.Len()", t)
	checkTrues("SymmetricDifferenfce{0,bitSize/2,bitSize}&{1,bitSize/2,bitSize+1}.Has()", t,
		data.Has(0),
		data.Has(1),
		data.Has(bitSize+1),
		data.Has(bitSize))
}

func TestElems(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	words := data.Elems()
	words2 := []int{0, bitSize, bitSize / 2}
	if !reflect.DeepEqual(words, words) {
		t.Errorf("actual: %v  expected: %v", words, words2)
	}
}

func TestRemove(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	data.Remove(bitSize / 2)
	checkInt(2, data.Len(), "Removed{0,bitSize}", t)
	checkTrues("Removed{0,bitSize}", t,
		data.Has(0),
		data.Has(bitSize))
}

func TestCopy(t *testing.T) {
	data := IntSet{}
	data.Add(0)
	data.Add(bitSize)
	data.Add(bitSize / 2)
	copied := data.Copy()
	if !reflect.DeepEqual(data.words, copied.words) {
		t.Errorf("actual: %v  expected: %v", data.String(), copied.String())
	}
	data.Add(1)
	if reflect.DeepEqual(data.words, copied.words) {
		t.Errorf("not equal actual: %v expected: %v", data.words, copied.words)
	}
}

func TestAddAll(t *testing.T) {
	data := IntSet{}
	data.AddAll(0, bitSize, bitSize/2)
	checkInt(3, data.Len(), "AddAll{0,bitSize/2,bitSize}", t)
	checkTrues("AddAll{0,bitSize/2bitSize}", t,
		data.Has(0),
		data.Has(bitSize/2),
		data.Has(bitSize))
}

func checkInt(expected, actual int, message string, t *testing.T) {
	if actual != expected {
		t.Errorf("%s checkInt Error expected: %d, actual: %d", message, expected, actual)
	}
}

func checkString(expected, actual string, message string, t *testing.T) {
	if actual != expected {
		t.Errorf("%s checkString Error expected: %s, actual: %s", message, expected, actual)
	}
}

func checkUint(expected, actual uint, message string, t *testing.T) {
	if actual != expected {
		t.Errorf("%s checkUint Error expected: %d, actual: %d", message, expected, actual)
	}
}

func checkFalse(expected bool, message string, t *testing.T) {
	if expected {
		t.Errorf("%s checkFalse Error", message)
	}
}

func checkTrue(expected bool, message string, t *testing.T) {
	if !expected {
		t.Errorf("%s checkTrue Error", message)
	}
}

func checkTrues(message string, t *testing.T, expected ...bool) {
	for _, exp := range expected {
		checkTrue(exp, message, t)
	}
}
