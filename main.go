package main

import (
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	"os"
	"io"
	"encoding/json"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"sort"
)
/*
type task interface {
	T1
	T2
}

type T1 struct {
	w, h int32
	s string
}

type T2 struct {
	envA, envB task2.Envelope
}
*/
const path = "./tasks.json"
/*
func main () {
	fmt.Println(math.Pow(2, 2), path)
}
*/

func main() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	var textCorr []byte
	for k, v := range text {
		if v == 0 {
			textCorr = text[0:k]
			break
		}
	}
	var dat map[string]map[string]string
	if err := json.Unmarshal(textCorr, &dat); err != nil {
		panic(err)
	}
	for k, v := range dat {
		for kTask, vTask := range v {
			switch k {
			case "task1":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				w, h, s, err := task1Validator(vTask)
				if err == nil {
					fmt.Println(task1.Task1(w, h, s))
				} else {
					checkError(err)
				}
			case "task2":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				env1AB, env1CD, env2AB, env2CD, err := task2Validator(vTask)
				if err == nil {
					fmt.Println(task2.Task2(task2.Envelope{env1AB, env1CD}, task2.Envelope{env2AB, env2CD}))
				} else {
					checkError(err)
				}
			case "task3":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				inp, err := task3Validator(vTask)
				if err == nil {
					fmt.Println(task3.Task3(inp))
				} else {
					checkError(err)
				}
			case "task4":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				in, err := task4Validator(vTask)
				if err == nil {
					poli, err := task4.Task4(in)
					if err == nil {
						fmt.Println(poli)
					} else {
						checkError(err)
					}
				} else {
					checkError(err)
				}
			case "task5":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				min, max, err := task5Validator(vTask)
				if err == nil {
					fmt.Println(task5.Task5(min, max))
				} else {
					checkError(err)
				}
			case "task6":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				length, pow, err := task6Validator(vTask)
				if err == nil {
					fmt.Println(task6.Task6(length, pow))
				} else {
					checkError(err)
				}
			case "task7":
				fmt.Println("\n", strings.ToUpper(k), "-", strings.ToTitle(kTask))
				err := task7Validator()
				if err == nil {
					fmt.Println(task7.Task7())
				} else {
					checkError(err)
				}
			}
		}
	}
}

func task7Validator() (error) {
	var file, err = os.OpenFile("context", os.O_RDWR, 0644)
	if err != nil {
		return errors.New("Problem with opening 'context' file.")
	}
	var text = make([]byte, 128)
	for {
		n, err := file.Read(text)
		if err != nil && err != io.EOF {
			return errors.New("Problem with reading 'context' file.")
		}
		if n == 0 {
			break
		}
	}
	var inS string
	for k, v := range text {
		if v == 0 {
			inS = string(text[0:k])
			break
		}
	}
	inI := strings.Split(inS, " ")
	_, err = strconv.Atoi(inI[0])
	if err != nil {
		return errors.New("Incorrect input. Expect int.")
	}
	if len(inI) == 2{
		_, err = strconv.Atoi(inI[1])
		if err != nil {
			return errors.New("Incorrect input. Expect int and int.")
		}
	}
	return nil
}

func task6Validator(in string) (int, int, error) {
	in = strings.Replace(in, " ", "", -1)
	val := strings.Split(in, ",")
	if len(val) != 2 {
		return 0, 0, errors.New("Incorrect input. Expect int and int.")
	}
	length, err := strconv.Atoi(val[0])
	if err != nil {
		return 0, 0, errors.New("Incorrect input. Expect int and int.")
	}
	pow, err := strconv.Atoi(val[1])
	if err != nil {
		return 0, 0, errors.New("Incorrect input. Expect int and int.")
	}
	return length, pow, nil
}

func task5Validator(in string) (int, int, error) {
	in = strings.Replace(in, " ", "", -1)
	val := strings.Split(in, ",")
	if len(val) != 2 {
		return 0, 0, errors.New("Incorrect input. Expect int and int.")
	}
	min, err := strconv.Atoi(val[0])
	if err != nil {
		return 0, 0, errors.New("Incorrect input. Expect int and int.")
	}
	max, err := strconv.Atoi(val[1])
	if err != nil {
		return 0, 0, errors.New("Incorrect input. Expect int and int.")
	}
	if min > max {
		return 0, 0, errors.New("Incorrect input. Input should be first < second")
	}
	return min, max, nil
}

func task4Validator(in string) (int, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return 0, errors.New("Incorrect input. Must be int64.")
	}
	if out < 10 {
		return 0, errors.New("Incorrect input. Input must be greater than 0. ")
	}
	return out, nil
}

func task3Validator(in string) ([]task3.Triangle, error) {
	var out []task3.Triangle
	in = strings.Replace(in, " ", "", -1)
	in = strings.Replace(in, "task3.Triangle{", "", -1)
	in = strings.Replace(in, "'", "", -1)
	inA := strings.Split(in, "},")
	for _, v := range inA {
		v = strings.Replace(v, "}", "", -1)
		val := strings.Split(v, ",")
		if len(val) != 4 {
			return []task3.Triangle{}, errors.New("Wrong input. Triangle has only name and three angles. ")
		}
		var valF []float32
		for i := 1; i < len(val); i++ {
			tmp, err := strconv.ParseFloat(val[i], 32)
			if err != nil {
				return []task3.Triangle{}, errors.New("Wrong input. Angles must be float32. ")
			}
			valF = append(valF, float32(tmp))
		}
		sort.Slice(valF, func(i, j int) bool { return valF[i] < valF[j] })
		if (valF[0] + valF[1]) < valF[2] {
			return []task3.Triangle{}, errors.New("Wrong input. " + val[0] + " : not a triangle. Check input. ")
		}
		out = append(out, task3.Triangle{val[0], valF[0], valF[1], valF[2]})
	}
	return out, nil
}

func task2Validator(in string) (float32, float32, float32, float32, error) {
	var eee []float32
	in = strings.Replace(in, " ", "", -1)
	in = strings.Replace(in, "task2.Envelope{", "", -1)
	in = strings.Replace(in, "}", "", -1)
	inA := strings.Split(in, ",")
	if len(inA) > 4 {
		return 0.0, 0.0, 0.0, 0.0, errors.New("According to rules you should put only 2 Envelopes.")
	}
	for _, v := range inA {
		ee, err := strconv.ParseFloat(v, 32)
		if err != nil {
			checkError(err)
		} else {
			eee = append(eee, float32(ee))
		}
	}
	if eee[0] <= 0 || eee[1] <= 0 || eee[2] <= 0 || eee[3] <= 0 {
		return 0.0, 0.0, 0.0, 0.0, errors.New("Every size of Envelope must be > 0.")
	}
	return eee[0], eee[1], eee[2], eee[3], nil
}

func task1Validator(in string) (int, int, string, error) {
	in = strings.Replace(in, " ", "", -1)
	inW := strings.Split(in, ",")
	w, err := strconv.Atoi(inW[0])
	if err != nil {
		return 0, 0, "", errors.New("Incorrect input. Expect int, int, symbol and coma as a separator.")
	}
	h, err := strconv.Atoi(inW[1])
	if err != nil {
		return 0, 0, "", errors.New("Incorrect input. Expect int, int, symbol and coma as a separator.")
	}
	s := inW[2]
	if w <= 0 || h <= 0 || s == "" {
		return 0, 0, "", errors.New("Incorrect input. Width and height must be greater than 0 and symbol cant be nil.")
	}
	return w, h, s, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
