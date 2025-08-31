package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// I dont not generally comment like this while coding but if someone reads it i tried to explain like i will explain in in-person
// Here i am calculating the power of 4
func powOf4(x int) int {
	return x * x * x * x
}

// according to the problem statement we will be calculating sum of (Yn^4)
func hennge(lines []string, index, total int, res []int) []int {
	//This is the base case for recursion as we are not allowed to use For loop
	if index >= len(lines) || len(res) == total {
		return res
	}
	X, error := strconv.Atoi(lines[index]) // over here i am converting current line to integer
	switch {
	case error != nil, X <= 0, X > 100: // checking for errors and constraints as specified in the problem statement
		res = append(res, -1)
		return res
	}
	index++
	if index >= len(lines) { //checking if index is out of bounds
		res = append(res, -1)
		return res
	}
	nums := strings.Fields(lines[index]) // over here i am splitting the line into fields
	switch {                             //switch case for checking if number of fields is equal to X or not
	case len(nums) != X:
		res = append(res, -1)
		return res
	}
	sum, valid := sumNegPowOf4Recursive(nums, 0)
	if !valid {
		return hennge(lines, index+1, total, append(res, -1))
	}
	return hennge(lines, index+1, total, append(res, sum))
	// Okay, so here I need to process each test case recursively.
	// First, I'll check if I've processed enough test cases or run out of input lines.
	// If so, I can just return the results I've collected so far.
}

// sumNegPowOf4Recursive recursively sums powOf4 for non-positive numbers in nums.
// Returns (sum, true) if all numbers are valid, or (0, false) if any conversion fails.
func sumNegPowOf4Recursive(nums []string, sum int) (int, bool) {
	if len(nums) == 0 {
		return sum, true
	}
	n, error := strconv.Atoi(nums[0])
	switch {
	case error != nil:
		return 0, false
	case n <= 0:
		return sumNegPowOf4Recursive(nums[1:], sum+powOf4(n))
	default:
		return sumNegPowOf4Recursive(nums[1:], sum)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) // buffered reader so i can read input from standard input
	line, error := reader.ReadString('\n')
	if error != nil { //incase of error exit
		return
	}
	line = strings.TrimSpace(line) //remove any trailing or leading whitespace
	N, error := strconv.Atoi(line)
	if error != nil || N < 1 || N > 100 {
		return
	}
	inputLines := readInputLines(reader, N, []string{})
	res := hennge(inputLines, 0, N, []int{})
	printResults(res, 0)
}

// readInputLines reads input lines from the buffered reader
// this function returns a slice of strings containing the input lines after all test cases are read
func readInputLines(reader *bufio.Reader, n int, str []string) []string {
	if n == 0 {
		return str
	}
	lineX, error := reader.ReadString('\n')
	if error != nil {
		str = append(str, "")
		str = append(str, "")
		return readInputLines(reader, n-1, str)
	}
	str = append(str, strings.TrimSpace(lineX))
	lineNums, error := reader.ReadString('\n')
	if error != nil {
		str = append(str, "")
		return readInputLines(reader, n-1, str)
	}
	str = append(str, strings.TrimSpace(lineNums))
	return readInputLines(reader, n-1, str)
}

func printResults(res []int, index int) {
	if index >= len(res) { //base case for recursion
		return
	} //printing results and ensuring proper formatting like no extra blanks at the start or end
	if index > 0 {
		fmt.Print("\n")
	}
	fmt.Print(res[index])
	printResults(res, index+1)
}
