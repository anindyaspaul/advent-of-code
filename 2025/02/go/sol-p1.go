package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func powerOfTen(e int) int {
	res := 1
	for e > 0 {
		res *= 10
		e--
	}
	return res
}

func repeatedNines(count int) int {
	res := 0
	for count > 0 {
		res *= 10
		res += 9
		count--
	}
	return res
}

func main() {
	log.SetFlags(0) // Disable date and time logging
	if _, isDebug := os.LookupEnv("DEBUG"); !isDebug {
		log.SetOutput(io.Discard)
	}

	if len(os.Args) != 2 {
		log.Fatal("Error: Usage: go run sol-p1.go file")
	}

	// Precompute all invalid ids and their cumulative sums
	var invalidIds [100000]int
	var invalidIdsCumSum [100000]int
	idx := 1
	nDigits := 1
	for nDigits <= 5 {
		lower := powerOfTen(nDigits - 1)
		upper := repeatedNines(nDigits)

		for lower <= upper {
			// Create invalid ids by concatenating two numbers
			invalidIds[idx] = lower * powerOfTen(nDigits) + lower
			invalidIdsCumSum[idx] = invalidIdsCumSum[idx-1] + invalidIds[idx]
			lower++
			idx++
		}

		nDigits++
	}
	log.Println("Count of invalid IDs:", idx-1)
	

	fileName := os.Args[1]
	fmt.Println("Input file:", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	idRangesText := scanner.Text()
	for idRange := range strings.SplitSeq(idRangesText, ",") {
		// Parse each range
		log.Println(idRange)
		bounds := strings.Split(idRange, "-")
		lowerStr := bounds[0]
		upperStr := bounds[1]
		lower, _ := strconv.Atoi(lowerStr)
		upper, _ := strconv.Atoi(upperStr)
		log.Println(lower, upper)

		// Find the lower and upper bounds in the invalid numbers array
		lowerIdx := sort.SearchInts(invalidIds[:], lower)
		upperIdx := sort.SearchInts(invalidIds[:], upper)
		if upperIdx == len(invalidIds) || invalidIds[upperIdx] != upper {
			upperIdx--
		}
		log.Println(lowerIdx, invalidIds[lowerIdx], upperIdx, invalidIds[upperIdx])
		sum += invalidIdsCumSum[upperIdx] - invalidIdsCumSum[lowerIdx-1]
	}

	file.Close()

	fmt.Println("Sum of invalid ids:", sum)
}

