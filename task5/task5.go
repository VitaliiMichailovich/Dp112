package task5

import (
	"fmt"
	"strconv"
	"github.com/VitaliiMichailovich/DP112/taskregister"
	"encoding/json"
)

type Params struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Ticket struct {
	method    string
	easy      int
	difficult int
}

func Task(bytesParams []byte) (string, error) {
	var param Params
	err := json.Unmarshal(bytesParams, &param)
	if err != nil {
		return "", err
	}
	ticket, err := task5Validator(param.Min, param.Max)
	ticketReturn := ticket.method+", "+strconv.Itoa(ticket.easy)+", "+strconv.Itoa(ticket.difficult)
	return ticketReturn, err
}

func init() {
	taskregister.InitializeTask(5, Task)
}

func isHappyTicket(ticket int) (easy, difficult int) {
	easy, difficult = 0, 0
	var ticketArr [6]int
	var odd, even int
	for i := 5; i >= 0; i-- {
		ticketArr[i] = ticket % 10
		ticket = ticket / 10
	}
	if ticketArr[0]+ticketArr[1]+ticketArr[2] == ticketArr[3]+ticketArr[4]+ticketArr[5] {
		easy = 1
	}
	for _, i := range ticketArr {
		if i == 0 {

		} else if i%2 == 0 {
			even = even + i
		} else {
			odd = odd + i
		}
	}
	if odd == even {
		difficult = 1
	}
	return easy, difficult
}

func doTask5(min, max int) Ticket {
	var out Ticket
	var easy, difficult int
	for i := min; i <= max; i++ {
		easyFunc, difficultFunc := isHappyTicket(i)
		easy = easy + easyFunc
		difficult = difficult + difficultFunc
	}
	if easy > difficult {
		out = Ticket{"easy", easy, difficult}
	} else {
		out = Ticket{"difficult", easy, difficult}
	}
	return out
}

func task5Validator(min, max interface{}) (Ticket, error) {
	minInt, ok := min.(int)
	if !ok {
		minI, err := strconv.Atoi(min.(string))
		if err != nil {
			return Ticket{}, fmt.Errorf("Incorrect input: \"%v\". Must be int64. \n%v", min, err)
		}
		minInt = minI
	}
	maxInt, ok := max.(int)
	if !ok {
		maxI, err := strconv.Atoi(max.(string))
		if err != nil {
			return Ticket{}, fmt.Errorf("Incorrect input: \"%v\". Must be int64. \n%v", max, err)
		}
		maxInt = maxI
	}
	if minInt < 0 || maxInt <= 0 {
		return Ticket{}, fmt.Errorf("Incorrect input. Minimal ticket number (%v) must be >= 0 and Maximal (%v) ticket must be >0. ", min, max)
	}
	if minInt > maxInt {
		return Ticket{}, fmt.Errorf("Incorrect input. Minimal ticket number (%v) must be <= Maximal (%v) ticket. ", min, max)
	}
	if (minInt/1000000 > 0) || (maxInt/1000000 > 0) {
		return Ticket{}, fmt.Errorf("Incorrect input. To big tickets (%v, %v). Ticket vust be < 999999.", min, max)
	}
	return doTask5(minInt, maxInt), nil
}
