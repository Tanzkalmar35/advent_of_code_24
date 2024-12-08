package main

import (
	"bufio"
	"os"
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

func isSafe(input []int) bool {
	lastNum := -1
	// Representing the direction we're going. Either "increasing" or "decreasing".
	var direction string

	for idx, num := range input {
		if idx == 0 {
			lastNum = num
			continue
		} else if idx == 1 {
			if lastNum < num {
				direction = "increasing"
			} else if lastNum > num {
				direction = "decreasing"
			} else {
				return false
			}
		}

		numDif := 0

		if direction == "increasing" {
			numDif = num - lastNum
			if lastNum >= num {
				return false
			}
		} else if direction == "decreasing" {
			numDif = lastNum - num
			if lastNum <= num {
				return false
			}
		} else {
			panic("That was not supposed to happen...")
		}

		if numDif < 1 || numDif > 3 {
			return false
		}

		lastNum = num
	}

	return true
}

func LineReportIsSafe(line string) bool {
	safetyLevelsStrings := strings.Split(line, " ")
	var safetyLevels []int

	// Converting the strings to ints
	for _, numberAsString := range safetyLevelsStrings {
		number, err := strconv.Atoi(numberAsString)

		if err != nil {
			panic("Invalid puzzle input")
		}

		safetyLevels = append(safetyLevels, number)
	}

	return isSafe(safetyLevels)
}

func main() {
	filename := "input.txt"
	fileContent := ReadFile(filename)
	amountOfSafeReports := 0

	for _, line := range fileContent {
		if LineReportIsSafe(line) {
			amountOfSafeReports++
		}
	}

	println(amountOfSafeReports)

}
