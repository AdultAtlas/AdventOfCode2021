package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println("Hey hey we got an error!")
	}

	defer file.Close()

	var lines []string
	var previousNum int = -1
	var ascendingCount int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, element := range lines {
		value, err := strconv.Atoi(element)

		if previousNum >= 0 && value > previousNum && err == nil {
			ascendingCount = ascendingCount + 1
		}
		previousNum = value
	}

	fmt.Println(ascendingCount)

}
