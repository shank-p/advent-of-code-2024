/*
	--- Day 2: Red-Nosed Reports ---
	https://adventofcode.com/2024/day/2
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)


func reportIsSafe(report []int) bool {
	for i:=0; i<len(report)-2; i++ {
		if !(((report[i]<report[i+1] && report[i+1]<report[i+2]) || (report[i]>report[i+1] && report[i+1]>report[i+2])) &&
		(math.Abs(float64(report[i]-report[i+1]))<=3 && math.Abs(float64(report[i+1]-report[i+2]))<=3)) {
			return false
		}
	}
	return true
}

func partOne(data [][]int) (safeReports int) {
	for _, report := range data {
		reportLen := len(report)
		if reportLen<3 {
			safeReports+=1
			continue
		}
		if isSafe := reportIsSafe(report); isSafe == true {
			safeReports+=1
		}

	}
	return safeReports
}

func partTwo(data [][]int) (safeReports int) {
	for _, report := range data {
		reportLen := len(report)
		if reportLen<3 {
			safeReports+=1
			continue
		}
		for i:=0; i<reportLen; i++ {
			newSlice := slices.Concat(report[:i], report[i+1:])
			if isSafe := reportIsSafe(newSlice); isSafe == true {
				safeReports+=1
				break
			}
		}
	}
	return safeReports
}


func main() {
	file, err := os.Open("red_nosed_reports_input.txt")
	if err != nil {
		fmt.Println("# ===== Error reading file input! ===== #")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reportData [][]int
	for scanner.Scan() {
		report := []int {}
		items := strings.Fields(scanner.Text())
		for _, token := range items {
			val, _ := strconv.Atoi(token)
			report = append(report, val)
		}
		reportData = append(reportData, report)
	}

	fmt.Println("RES partOne ==>", partOne(reportData))
	fmt.Println("RES partTwo ==>", partTwo(reportData))
}

