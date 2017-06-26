package task1

func Task1(h int, w int, s string) string {
	var out string
	for i := 1; i <= w; i++ {
		for ii := 1; ii <= h; ii++ {
			if (ii+i)%2 != 0 {
				out += " "
			} else {
				out += s
			}
		}
		if i != w {
			out += "\n"
		}
	}
	return out
}
