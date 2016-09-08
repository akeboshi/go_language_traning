// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package sexpr

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	data := `
	((Title  "Dr. Strangelove")
                 (Subtitle  "How I Learned to Stop Worrying and Love the Bomb")
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
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	var movie Movie
	NewDecoder(bytes.NewBufferString(data)).Decode(&movie)
	t.Log(movie.Oscars)
}
