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
	file, err := os.Open("day02/" + filename)
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

func RemoveItem(list []int, idx int) []int {
	// for idx, num := range list {
	// 	println("List before contains " + strconv.Itoa(num) + " at " + strconv.Itoa(idx))
	// }
	// appended := append(list[:idx], list[idx+1:]...)
	// for idx, num := range appended {
	// 	println("List after contains " + strconv.Itoa(num) + " at " + strconv.Itoa(idx))
	// }
	return append(list[:idx], list[idx+1:]...)
}

// Part 1

// func isSafe(input []int) bool {
// 	lastNum := -1
// 	// Representing the direction we're going. Either "increasing" or "decreasing".
// 	var direction string
//
// 	for idx, num := range input {
// 		if idx == 0 {
// 			lastNum = num
// 			continue
// 		} else if idx == 1 {
// 			if lastNum < num {
// 				direction = "increasing"
// 			} else if lastNum > num {
// 				direction = "decreasing"
// 			} else {
// 				return false
// 			}
// 		}
//
// 		numDif := 0
//
// 		if direction == "increasing" {
// 			numDif = num - lastNum
// 			if lastNum >= num {
// 				return false
// 			}
// 		} else if direction == "decreasing" {
// 			numDif = lastNum - num
// 			if lastNum <= num {
// 				return false
// 			}
// 		} else {
// 			panic("That was not supposed to happen...")
// 		}
//
// 		if numDif < 1 || numDif > 3 {
// 			return false
// 		}
//
// 		lastNum = num
// 	}
//
// 	return true
// }

// Part 2

func abc(list []int) bool {
	return isSafe(list, false)
}

func isSafe(input []int, hasRetry bool) bool {
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
				if hasRetry {
					newList := RemoveItem(input, idx)
					return abc(newList)
				} else {
					println("False as number is equal to last one - Without retry")
					return false
				}
			}
		}

		numDif := 0

		if direction == "increasing" {
			numDif = num - lastNum
			if lastNum >= num {
				if hasRetry {
					newList := RemoveItem(input, idx)
					return abc(newList)
				} else {
					println("False as direction is not increasing where it should - Without retry")
					return false
				}
			}
		} else if direction == "decreasing" {
			numDif = lastNum - num
			if lastNum <= num {
				if hasRetry {
					newList := RemoveItem(input, idx)
					return abc(newList)
				} else {
					println("False as direction is not decreasing where it should - Without retry")
					return false
				}
			}
		} else {
			panic("That was not supposed to happen...")
		}

		if numDif < 1 || numDif > 3 {
			if hasRetry {
				newList := RemoveItem(input, idx)
				println("False as numDif does not match - With retry")
				return abc(newList)
			} else {
				println("False as numDif does not match - Without retry")
				return false
			}
		}

		lastNum = num
	}

	return true
}

func LineReportIsSafe(line string) bool {
	safetyLevelsStrings := strings.Split(line, " ")
	var safetyLevels []int

	// Converting the strings to ints
	println("Member of list: ")
	for _, numberAsString := range safetyLevelsStrings {
		number, err := strconv.Atoi(numberAsString)

		if err != nil {
			panic("Invalid puzzle input")
		}

		safetyLevels = append(safetyLevels, number)
		print(strconv.Itoa(number))
	}

	println("")
	return isSafe(safetyLevels, true)
}

func main() {
	filename := "input.txt"
	fileContent := ReadFile(filename)
	amountOfSafeReports := 0

	for _, line := range fileContent {
		if LineReportIsSafe(line) {
			println("Above line is safe")
			amountOfSafeReports++
		}
	}

	println(amountOfSafeReports)
}
