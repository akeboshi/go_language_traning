package main

import "fmt"

func main() {
	/*
		s := "123-421394219-321"
		println(regexp.MustCompile(`^[0-9-]+$`).MatchString(s))
	*/ /*
		type S []S
		var s S
		s = append(s, s)
		fmt.Println(len(s))
		fmt.Printf("%p %p %p\n", s, s[0], s[0][0])
		fmt.Println(s)
	*/
	s := []string{}
	s = append(s, "ear")
	fmt.Printf("%p %p\n", s, &s[0])
}

/*
	type Movie struct {
		Title, Subtitle string
		Year            int
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
	a, _ := json.Marshal(strangelove)
	fmt.Println(string(a))
}*/