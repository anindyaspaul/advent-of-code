package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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

func countDigits(x int) int {
	count := 0
	for x > 0 {
		count++
		x /= 10
	}
	return count
}

func repeat(x int, count int) int {
	result := 0
	nDigits := countDigits(x)
	for ; count > 0; count-- {
		result = result*powerOfTen(nDigits) + x
	}
	return result
}

func main() {
	log.SetFlags(0) // Disable date and time logging
	if _, isDebug := os.LookupEnv("DEBUG"); !isDebug {
		log.SetOutput(io.Discard)
	}

	if len(os.Args) != 2 {
		log.Fatal("Error: Usage: go run sol-p1.go file")
	}

	// Loop through 1 to 100000 and create invalid numbers by repeating them
	invalidIdsSet := make(map[int]bool)
	for i := 1; i < 100000; i++ {
		if _, ok := invalidIdsSet[i]; ok {
			continue
		}
		nDigits := countDigits(i)
		for repeatCount := 2; repeatCount*nDigits <= 10; repeatCount++ {
			x := repeat(i, repeatCount)
			invalidIdsSet[x] = true
		}
	}

	log.Println("Total invalids:", len(invalidIdsSet))

	invalidIds := make([]int, len(invalidIdsSet)+1)
	invalidIdsCumSum := make([]int, len(invalidIdsSet)+1)

	// Take the invalid ids from the set
	i := 1
	for k := range invalidIdsSet {
		invalidIds[i] = k
		i++
	}

	// Sort the set for binary search and to calculate cumulative sums
	slices.Sort(invalidIds)
	for i = 1; i < len(invalidIds); i++ {
		invalidIdsCumSum[i] = invalidIdsCumSum[i-1] + invalidIds[i]
	}

	log.Println("Invalid ids: ", invalidIds[:50])
	log.Println("Cumsum: ", invalidIdsCumSum[:50])

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
