package task4

import (
	"errors"
	"fmt"
	"strconv"
	"sort"
)

func palindromeChecker(in []int64) bool {
	var out bool = true
	var f int = len(in) / 2
	for i := 0; i < f; i++ {
		if in[i] != in[len(in)-1-i] {
			out = false
		}
	}
	return out
}

func doTask4(in int64) (int64, error) {
	var out int64
	var err error
	var nums []int64
	var poli []int64
	for {
		nums = append(nums, int64(in%10))
		in = in / 10
		if in == 0 {
			break
		}
	}
	for i := len(nums); i > 1; i-- {
		for ii := 0; ii < len(nums)-i+1; ii++ {
			if palindromeChecker(nums[ii : ii+i]) {
				var tmp int64
				for iii := ii; iii < ii+i; iii++ {
					tmp = tmp*10 + nums[iii]
				}
				poli = append(poli, tmp)
			}
		}
		if len(poli) > 0 {
			sort.Slice(poli, func(i, j int) bool { return poli[i] > poli[j] })
			out = poli[0]
		}
	}
	if out == 0 {
		err = errors.New("There is NO Palindrome. ")
	}
	return out, err
}

func Task4(in interface{}) (int64, error) {
	number, ok := in.(int64)
	if !ok {
		numberReturn, err := strconv.ParseInt(in.(string), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("Incorrect input: \"%v\". Must be int64. \n%v", in, err)
		}
		number = numberReturn
	}
	if number <= 10 {
		return 0, fmt.Errorf("Incorrect input \"%v\". Input must be greater than 10.", in)
	}
	returnTask4, err := doTask4(number)
	if err != nil {
		return 0, err
	}
	return returnTask4, nil
}