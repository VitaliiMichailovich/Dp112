package main

import (
	"../src/task1"
	"../src/task2"
	"../src/task3"
	"../src/task4"
	"../src/task5"
	"../src/task6"
	"fmt"
)

func main() {
	fmt.Println(task1.Task1(12, 7, "*"))
	fmt.Println(task2.Task2(task2.Envelope{20.0, 30.0}, task2.Envelope{10.0, 5.0}))
	fmt.Println(task3.Task3(task3.Triangle{"j1", 10, 20, 725}, task3.Triangle{"j2", 60, 22.5, 77.8}, task3.Triangle{"j3", 22.2, 20.5, 30.8}, task3.Triangle{"j4", 1.1, 2.71, 3.4}, task3.Triangle{"j5", 0.18, 0.20, 0.30}, task3.Triangle{"j6", 270, 270, 3.0}, task3.Triangle{"j7", 1750, 1750.20, 10.30}))
	fmt.Println(task4.Task4(12345746437))
	fmt.Println(task5.Task5(75319, 100952))
	task6.Task6(7, 777)
}
