/*
	--- Day 1: Historian Hysteria ---
	-> https://adventofcode.com/2024/day/1
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne(leftList []int, rightList []int) (sumDiff uint) {
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})
	for i:=0; i<len(leftList);i++ {
		sumDiff += uint(math.Abs(float64(leftList[i])-float64(rightList[i])))
	}
	return sumDiff
}

func partTwo(leftList []int, rightList []int) (similarityScore int) {
	rightListMap := make(map[int]uint)
	for _, num := range rightList {
		if _, ok := rightListMap[num]; ok == true {
			rightListMap[num]+=1
		} else {
			rightListMap[num]=1
		}
	}

	for _, num := range leftList {
		if val, ok := rightListMap[num]; ok == true {
			similarityScore+= (num*int(val))
		}
	}

	return similarityScore
}

func main() {
	file, err := os.Open("historian_hysteria_input.txt")
	if err != nil {
		fmt.Println("# ===== Error reading file input! ===== #")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftList, rightList []int
	var leftNum, rightNum int
	for scanner.Scan() {
		numSplit := strings.Fields(scanner.Text())
		leftNum, _ = strconv.Atoi(numSplit[0])
		rightNum, _ = strconv.Atoi(numSplit[1])
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	
	fmt.Println("RES partOne => ", partOne(leftList, rightList))
	fmt.Println("RES partTwo => ", partTwo(leftList, rightList))
}