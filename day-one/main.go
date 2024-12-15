package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Time to try advent of code")

	data := readFile("input.txt")

	defer data.Close()
	if data == nil {
		return
	}

	leftSide := []int{}
	rightSide := []int{}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "  ")
		leftNum, err := strconv.Atoi(strings.TrimSpace(splitLine[0]))

		if err != nil {
			fmt.Println("Invalid number:", err)
			return
		}
		rightNum, err := strconv.Atoi(strings.TrimSpace(splitLine[1]))

		if err != nil {
			fmt.Println("Invalid number:", err)
			return
		}

		leftSide = append(leftSide, leftNum)
		rightSide = append(rightSide, rightNum)
	}

	sort.Ints(leftSide)
	sort.Ints(rightSide)

	totalDistance := 0
	for index := 0; index < len(leftSide); index++ {
		_, _, distance := getPairDistance(leftSide, rightSide, index)

		totalDistance += distance
	}

	counts := make(map[int]int)

	for _, value := range rightSide {
		counts[value]++
	}

	similaritiesTotal := 0

	for i := 0; i < len(leftSide); i++ {
		leftSideValue := leftSide[i]

		if count, exists := counts[leftSideValue]; exists {
			similaritiesTotal += leftSideValue * count
		}
	}

	fmt.Println("Total distance is: ", totalDistance)
	fmt.Println("Total similarities is: ", similaritiesTotal)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}

func getPairDistance(leftSide, rightSide []int, index int) (int, int, int) {
	pairDistance := 0

	if rightSide[index] > leftSide[index] {
		pairDistance = rightSide[index] - leftSide[index]
	} else {
		pairDistance = leftSide[index] - rightSide[index]
	}

	return leftSide[index], rightSide[index], pairDistance

}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	return file
}
