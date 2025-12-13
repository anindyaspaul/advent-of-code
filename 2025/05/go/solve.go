package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(file *os.File) ([][2]int, []int) {
	isDbEntry := true
	var ranges [][2]int
	var queries []int
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isDbEntry = false
			continue
		}
		if isDbEntry {
			boundStrings := strings.Split(line, "-")
			lowerBound, _ := strconv.Atoi(boundStrings[0])
			upperBound, _ := strconv.Atoi(boundStrings[1])
			ranges = append(ranges, [2]int{lowerBound, upperBound})
		} else {
			queryId, _ := strconv.Atoi(line)
			queries = append(queries, queryId)
		}
	}
	return ranges, queries
}

func solvePart1Naive(ranges [][2]int, queries []int) {
	count := 0
	for _, id := range queries {
		for _, bounds := range ranges {
			if id >= bounds[0] && id <= bounds[1] {
				count++
				break
			}
		}
	} 
	fmt.Println("Solution to part 1:", count)
}

func solvePart2(ranges [][2]int) {
	log.Println(ranges)
	slices.SortFunc(ranges, func(a [2]int, b [2]int) int {
		if b[0] == a[0] {
			return a[1] - b[1]
		} else {
			return a[0] - b[0]
		}
	})
	log.Println(ranges)

	count := 0
	window := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] > window[1] {
			count += window[1] - window[0] + 1
			window = ranges[i]
		} else {
			if ranges[i][1] > window[1] {
				window[1] = ranges[i][1]
			}
		}
	}
	count += window[1] - window[0] + 1
	fmt.Println("Solution to part 2:", count)
}

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalln("usage: go run solve.go input.txt")
	}

	if _, debug := os.LookupEnv("DEBUG"); !debug {
		log.SetOutput(io.Discard)
	}

	fileName := os.Args[1]
	fmt.Println("Input file:", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	ranges, queries := parseInput(file)

	solvePart1Naive(ranges, queries)
	solvePart2(ranges)

	file.Close()
}
