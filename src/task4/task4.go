package task4

import "errors"

func Task4(in int) (int, error) {
	var out int
	var err error
	var nums []int
	var poli []int
	for {
		nums = append(nums, in%10)
		in = in / 10
		if in == 0 {
			break
		}
	}
	for i := len(nums); i >= 1; i-- {
		for ii := 0; ii < len(nums)-i+1; ii++ {
			if poliCheck(nums[ii : ii+i]) {
				poli = nums[ii : ii+i]
				break
			}
		}
		if len(poli) > 1 {
			for ii := 0; ii < len(poli); ii++ {
				out = out*10 + poli[ii]
			}
			break
		}
	}
	if out == 0 {
		err = errors.New("There is NO Palindrome. ")
	}
	return out, err
}

func poliCheck(in []int) bool {
	var out bool = true
	var f int = len(in) / 2
	for i := 0; i < f; i++ {
		if in[i] != in[len(in)-1-i] {
			out = false
		}
	}
	return out
}
