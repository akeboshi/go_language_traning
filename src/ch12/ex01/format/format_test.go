// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package format

import (
	"reflect"
	_ "testing"
)

// TODO: Unordered output が go 1.7 で実装される。Atom editorの都合上1.7が使えない
func ExampleDisplayMapWithArrayKey() {
	data := map[[2]string]string{[2]string{"foo"}: "bar"}
	display("data", reflect.ValueOf(data), 0)
	// Output:
	// data[[2]string{"foo", ""}] = "bar"
}

func ExampleDisplayMapWithStructKey() {
	type test struct {
		foo string
		bar string
	}
	data := map[test]string{test{"foo", "bar"}: "hoge"}
	display("data", reflect.ValueOf(data), 0)
	// Output:
	// data[format.test: {foo: "foo", bar: "bar"}] = "hoge"
}

// TODO: Unordered output が go 1.7 で実装される。Atom editorの都合上1.7が使えない
func ExampleDisplayCycle() {
	type test struct {
		foo string
		bar *test
	}
	var data test
	data = test{"foo", &data}
	display("data", reflect.ValueOf(data), 0)
	// Output:
	// data.foo = "foo"
	// (*data.bar).foo = "foo"
	// (*(*data.bar).bar).foo = "foo"
}

/*
  Mapの出力順が不定のため、表示のテストにExampleが使えない
  Panicにならないか位のチェックになっている。
*/
func ExampleDisplay() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	data := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("data", data)
}
