// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package sexpr

import (
	"bytes"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	data := `
	((Title  "Dr. Strangelove")
                 (Subtitle  "How I Learned to Stop Worrying and Love the Bomb")
								 (TBool t)
								 (ttag t)
								 (TFloat 1.234)
                 (Year  1964)
                 (Actor  (("Dr. Strangelove" "Peter Sellers")
                         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
                         ("Pres. Merkin Muffley" "Peter Sellers")
                         ("Gen. Buck Turgidson" "George C. Scott")
                         ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
                         ("Maj. T.J. \"King\" Kong" "Slim Pickens")))
                 (Oscars  ("Best Actor (Nomin.)"
                        "Best Adapted Screenplay (Nomin.)"
                        "Best Director (Nomin.)"
                        "Best Picture (Nomin.)"))
                 (Sequel  nil))
`
	type Movie struct {
		Title, Subtitle string
		TBool           bool
		TTag            bool `sexpr:"ttag"`
		TFloat          float64
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	var movie Movie
	NewDecoder(bytes.NewBufferString(data)).Decode(&movie)
	t.Log(movie.Oscars)
	t.Log(movie.TBool)
	t.Log(movie.TFloat)
	t.Log(movie.TTag)
}

func TestEncode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int `sexpr:"tagYear"`
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
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

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	if !strings.Contains(string(data), "tagYear") {
		t.Errorf("cant recognized tag: tagYear. data = %s", data)
	}
	t.Logf("Marshal() = %s\n", data)
}
