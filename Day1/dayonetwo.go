package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Window struct {
	id         int
	first_val  int
	second_val int
	third_val  int
}

func (w Window) Sum() int {
	return w.first_val + w.second_val + w.third_val
}

func main() {
	file, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println("Hey hey we got an error!")
	}

	defer file.Close()

	var lines []string
	var sums []int
	var ascendingCount int = 0
	var window_list []Window

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, scanner.Text())
		sums = append(sums, sum)
	}

	id := 0
	for index, _ := range sums {
		offset := len(sums) - index

		if offset > 2 {
			window_list = append(window_list, Window{id, sums[index], sums[index+1], sums[index+2]})
		} else if offset > 1 {
			window_list = append(window_list, Window{id, sums[index], sums[index+1], 0})
		} else if offset > 0 {
			window_list = append(window_list, Window{id, sums[index], 0, 0})
		}
		id = id + 1
	}

	for i := 1; i < len(window_list); i++ {
		if window_list[i].Sum() > window_list[i-1].Sum() {
			ascendingCount = ascendingCount + 1
		}
	}

	fmt.Print(id)
	fmt.Print(window_list)
	fmt.Print(ascendingCount)

}
