package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := readFile("test.txt")

	if file == nil {
		return
	}

	scanner := bufio.NewScanner(file)

	initialSafeCount := 0

	reports := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		reports = append(reports, splitLine)

	}

	convertedReports := [][]int{}

	for _, report := range reports {
		levels := []int{}
		for _, level := range report {
			levelValue := convertToInteger(level)
			levels = append(levels, levelValue)
		}

		convertedReports = append(convertedReports, levels)
	}

	cleanedValues := [][]int{}

	for _, report := range convertedReports {
		// we remove a level process the new line if safe now we count
		// if we reach end and not safe we maintain not safe

		isSafe, unsafe := processLine(report)
		if isSafe {
			initialSafeCount++
		}

		cleanedValues = append(cleanedValues, unsafe)
	}

	for _, cleanValue := range cleanedValues {

		if len(cleanValue) == 0 {
			continue
		}

		isSafe, _ := processLine(cleanValue)

		if isSafe {
			initialSafeCount++
		}

	}

	fmt.Println(cleanedValues)

	fmt.Println("Initial Safe reports: ", initialSafeCount)
	fmt.Println("Final Safe reports:")

}

func convertToInteger(value string) int {
	intValue, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil {
		fmt.Println("Could not convert string: ", err)
		panic(err)
	}

	return intValue
}

func processLine(line []int) (bool, []int) {
	unsafe := []int{}
	difference := 0
	decreasing := false
	firstItem := line[0]

	secondItem := line[1]

	if firstItem > secondItem {
		decreasing = true
	} else {
		decreasing = false
	}

	for i := 0; i < len(line)-1; i++ {

		lineValue := line[i]

		nextValue := line[i+1]

		if !decreasing && nextValue < lineValue {
			unsafe = append(unsafe, line[:i]...)
			unsafe = append(unsafe, line[i+1:]...)

			return false, unsafe
		}

		if decreasing && nextValue > lineValue {
			unsafe = append(unsafe, line[:i]...)
			unsafe = append(unsafe, line[i+1:]...)
			return false, unsafe
		}

		if decreasing {
			difference = lineValue - nextValue
		} else {
			difference = nextValue - lineValue
		}

		isDiffereceOkay := differRatio(difference)

		if !isDiffereceOkay {
			unsafe = append(unsafe, line[:i]...)
			unsafe = append(unsafe, line[i+1:]...)
			return false, unsafe
		}

	}

	return true, unsafe
}

func differRatio(item int) bool {
	if item < 1 {
		return false
	}

	if item > 3 {
		return false
	}

	return true
}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error Opening file: ", err)
		return nil
	}

	return file

}
