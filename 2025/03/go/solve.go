package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func digitCharToValue(digitChar uint8) int {
	return int(digitChar) - 48
}

func solvePart1(file *os.File) {
	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		
		msb := line[len(line)-2]
		lsb := line[len(line)-1]
		for i := len(line)-3; i >= 0; i-- {
			if line[i] >= msb {
				if msb > lsb {
					lsb = msb
				}
				msb = line[i]
			}
		}
		joltage := digitCharToValue(msb)*10 + digitCharToValue(lsb)
		log.Println(joltage)
		res += joltage
	}

	fmt.Println("Solution of part 1:", res)
}

func solvePart2(file *os.File) {
	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		
		batteries := make([]uint8, 12)
		pos := make([]int, 12)
		for i := range 12 {
			// find batteries[i] value
			batteries[i] = line[pos[i]]
			if i < 11 {
				pos[i+1] = pos[i]+1
			}
			for j := pos[i]+1; j <= len(line)-12+i; j++ {
				if line[j] > batteries[i] {
					batteries[i] = line[j]
					if i < 11 {
						pos[i+1] = j+1
					}
				}
			}
		}

		// Convert selected batteries to joltage value
		joltage := 0
		for i := range len(batteries) {
			joltage = joltage*10 + digitCharToValue(batteries[i])
		}
		log.Println(joltage)

		res += joltage
	}

	fmt.Println("Solution of part 2:", res)
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
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// solvePart1(file)
	solvePart2(file)
}
