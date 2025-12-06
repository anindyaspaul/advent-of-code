package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(0) // Disable date and time logging
	if _, isDebug := os.LookupEnv("DEBUG"); !isDebug {
		log.SetOutput(io.Discard)
	}

	if len(os.Args) != 2 {
		log.Fatal("Error: Usage: go run sol-p2.go file")
	}

	fileName := os.Args[1]
	fmt.Println("Input file:", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}

	curPosition := 50
	countZero := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		log.Println(instruction)
		
		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:]) 
		if err != nil {
			log.Fatal("Failed to parse distance:", instruction[1:])
		}
		log.Println(direction, distance)

		countZero += distance / 100
		distance %= 100

		if distance == 0 {
			continue
		}

		if direction == "L" {
			distance = -distance
		}
		
		newPosition := (curPosition + distance + 100) % 100
		log.Println(curPosition, distance, newPosition)

		if newPosition == 0 || (distance > 0 && newPosition < curPosition) || (distance < 0 && newPosition > curPosition && curPosition != 0) {
			log.Println("INC")
			countZero++
		}

		curPosition = newPosition
	}

	file.Close()

	fmt.Println("Password:", countZero)
}
