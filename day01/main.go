package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	var result []string

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() // Ensure the file is closed when we're done

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Read and print each line
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func SortAndCleanInput(input []string) ([]int, []int) {
	var leftList, rightList []int

	for _, line := range input {
		stringSplit := strings.Split(line, "   ")

		leftListAppend, leftAtoiErr := strconv.Atoi(strings.TrimSpace(stringSplit[0]))
		rightListAppend, rightAtoiErr := strconv.Atoi(strings.TrimSpace(stringSplit[1]))

		if leftAtoiErr != nil && rightAtoiErr != nil {
			panic("Invalid input set given!")
		}

		leftList = append(leftList, leftListAppend)
		rightList = append(rightList, rightListAppend)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func main() {
	filename := "input.txt"
	fileContent := ReadFile(filename)
	result := 0

	sortedLeftList, sortedRightList := SortAndCleanInput(fileContent)

	if len(sortedLeftList) != len(sortedRightList) {
		panic("Invalid input set given! - Left and right side have different lengths")
	}

	for i := 0; i < len(sortedLeftList); i++ {
		if sortedLeftList[i] < sortedRightList[i] {
			result += sortedRightList[i] - sortedLeftList[i]
		} else if sortedLeftList[i] > sortedRightList[i] {
			result += sortedLeftList[i] - sortedRightList[i]
		} else {
			// If the two numbers are equal, there's no need to do anything,
			// as we would only add 0 to it.
		}
	}

	println(result)
}
