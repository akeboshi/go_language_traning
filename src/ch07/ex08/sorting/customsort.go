package sorting

type CustomSort struct {
	T     []*Track
	Lessf []func(x, y *Track) int
}

var SortAlgs = map[string]func(x, y *Track) int{
	"Title": func(x, y *Track) int {
		if x.Title == y.Title {
			return 0
		} else if x.Title < y.Title {
			return 1
		} else {
			return -1
		}
	},
	"Artist": func(x, y *Track) int {
		if x.Artist == y.Artist {
			return 0
		} else if x.Artist < y.Artist {
			return 1
		} else {
			return -1
		}
	},
	"Year": func(x, y *Track) int {
		if x.Year == y.Year {
			return 0
		} else if x.Year < y.Year {
			return 1
		} else {
			return -1
		}
	},
	"Album": func(x, y *Track) int {
		if x.Album == y.Album {
			return 0
		} else if x.Album < y.Album {
			return 1
		} else {
			return -1
		}
	},
	"Length": func(x, y *Track) int {
		if x.Length == y.Length {
			return 0
		} else if x.Length < y.Length {
			return 1
		} else {
			return -1
		}
	},
}

func (x CustomSort) Len() int { return len(x.T) }
func (x CustomSort) Less(i, j int) bool {
	for _, f := range x.Lessf {
		less := f(x.T[i], x.T[j])
		if less != 0 {
			return less > 0
		}
	}
	return false
}
func (x CustomSort) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }
