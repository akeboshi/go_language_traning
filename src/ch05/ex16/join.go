package main

func main() {
	println(join(",", "foo", "bar", "hoge"))
}

func join(sep string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	if len(vals) == 1 {
		return vals[0]
	}

	joined := vals[0]
	for _, val := range vals[1:] {
		joined += sep + val
	}
	return joined
}
