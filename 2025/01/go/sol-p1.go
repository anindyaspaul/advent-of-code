package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var isDebug = os.Getenv("DEBUG") == "1"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error: Usage: go run sol-p1.go file")
	}

	fileName := os.Args[1]
	if isDebug { fmt.Println("Input file:", fileName) }

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}

	curPosition := 50
	countZero := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		if isDebug { fmt.Println(instruction) }
		
		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:]) 
		if err != nil {
			log.Fatal("Failed to parse distance:", instruction[1:])
		}
		if isDebug { fmt.Println(direction, distance) }

		distance %= 100
		if direction == "L" {
			distance = -distance
		}
		
		newPosition := (curPosition + distance + 100) % 100
		if isDebug { fmt.Println(curPosition, distance, newPosition) }

		if newPosition == 0 {
			countZero++
		}

		curPosition = newPosition
	}

	file.Close()

	fmt.Println("Password:", countZero)
}
