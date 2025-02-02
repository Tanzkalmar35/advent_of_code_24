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
	file, err := os.Open("day03/" + filename)
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

func main() {
	filename := "input.txt"
	fileContent := ReadFile(filename)
	res := 0
	for _, line := range fileContent {
		res += ProcessInput(line)
	}
	println("Result is: " + strconv.Itoa(res))
}

func ProcessInput(line string) int {
	res := 0

	for idx, char := range strings.Split(line, "") {
		split := strings.Split(line, "")
		overIter := idx + 1
		if char != "m" {
			continue
		}
		if split[overIter] != "u" {
			continue
		}
		overIter++
		if split[overIter] != "l" {
			continue
		}
		overIter++
		if split[overIter] != "(" {
			continue
		}
		overIter++
		if num1, iter1, err1 := ProcessMulNum(overIter, split, ","); err1 != nil {
			continue
		} else {
			overIter = iter1
			if num2, _, err2 := ProcessMulNum(overIter, split, ")"); err2 != nil {
				continue
			} else {
				res += num1 * num2
			}
		}

	}

	return res
}

func ProcessMulNum(overIter int, split []string, toChar string) (int, int, error) {
	fromIdx := overIter
	toIdx := 0
	for {
		if split[overIter] == toChar {
			break
		}
		overIter++
		toIdx = overIter
	}
	var res strings.Builder
	for i := fromIdx; i < toIdx; i++ {
		a := split[i]
		println(a)
		res.WriteString(split[i])
	}
	num, err := strconv.Atoi(res.String())
	if err != nil {
		return 0, 0, err
	}
	return num, overIter + 1, nil
}
