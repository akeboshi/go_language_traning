//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import "sort"

func main() {
	list := sort.StringSlice{"a", "b"}
	println(IsPalindrome(list))
	list = sort.StringSlice{"abc", "foo", "foo", "abc"}
	println(IsPalindrome(list))
}

func IsPalindrome(s sort.Interface) bool {
	len := s.Len()
	for i := 0; i < len/2; i++ {
		j := len - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
