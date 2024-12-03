/*
	--- Day 3: Mull It Over ---
	https://adventofcode.com/2024/day/3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func partOne(file *os.File) (res int) {
	expression := regexp.MustCompile(`mul\([-]?\d{1,3},[-]?\d{1,3}\)`)
	var matches []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, expression.FindAllString(line, -1)...)
	}
	
	var nums []string
	var numA, numB int
	for _, instruction := range matches {
		nums = strings.Split(instruction[4:len(instruction)-1], ",")
		numA, _ = strconv.Atoi(nums[0])
		numB, _ = strconv.Atoi(nums[1])
		res += numA*numB
	}
	return res
}

func partTwo(file *os.File) (res int) {
	expression := regexp.MustCompile(`mul\([-]?\d{1,3},[-]?\d{1,3}\)|do\(\)|don't\(\)`)
	
	var matches []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches = append(matches, expression.FindAllString(line, -1)...)
	}
	
	var nums []string
	var numA, numB int
	var skipInstruction = false
	for _, instruction := range matches {
		if instruction=="do()" {
			skipInstruction = false
			continue
		} else if instruction=="don't()" {
			skipInstruction = true
			continue
		} else if skipInstruction!=true {
			nums = strings.Split(instruction[4:len(instruction)-1], ",")
			numA, _ = strconv.Atoi(nums[0])
			numB, _ = strconv.Atoi(nums[1])
			res += numA*numB
		}	
	}
	return res
}

func main() {
	file, err := os.Open("mull_it_over_input.txt")
	if err != nil {
		panic("Unable to read input file!")
	}
	defer file.Close()

	// fmt.Println("RES partOne =>", partOne(file))
	fmt.Println("RES partTwo =>", partTwo(file))
}