package task1

func Task1(w int, h int, s string) string {
	var out string
	for i := 1; i <= h; i++ {
		for ii := 1; ii <= w; ii++ {
			if (ii+i)%2 != 0 {
				out += " "
			} else {
				out += s
			}
		}
		if i != h {
			out += "\n"
		}
	}
	return out
}
