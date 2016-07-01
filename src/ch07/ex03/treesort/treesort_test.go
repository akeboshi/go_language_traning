//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package treesort

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestTreeString(t *testing.T) {
	data := []int{5, 3, 6, 8, 10, 1}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}

	actual := fmt.Sprint(root)
	for _, d := range data {
		n := strconv.Itoa(d)
		check := strings.Contains(actual, n)
		if !check {
			t.Errorf("error expected:%s actual:%s", intSliceToStr(data), actual)
		}
	}
}

func intSliceToStr(sl []int) string {
	str := ""
	for _, s := range sl {
		str += strconv.Itoa(s) + " "
	}
	return str[:len(str)-1]
}
