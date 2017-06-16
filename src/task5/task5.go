package task5

type Ticket struct {
	method    string
	easy      int
	difficult int
}

func Task5(min, max int) Ticket {
	var out Ticket
	var easy, difficult int
	for i := min; i <= max; i++ {
		v := i
		var even, odd, half1, half2 int
		for ii := 0; ii != -1; ii++ {
			num := v % 10
			v = v / 10
			if ii%2 == 0 {
				even += num
			} else {
				odd += num
			}
			if ii <= 2 {
				half1 += num
			} else {
				half2 += num
			}
			if v == 0 {
				break
			}
		}
		if even == odd {
			difficult ++
		}
		if half1 == half2 {
			easy++
		}
	}
	if easy > difficult {
		out = Ticket{"easy", easy, difficult}
	} else {
		out = Ticket{"difficult", easy, difficult}
	}
	return out
}