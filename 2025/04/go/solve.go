package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var d = [][]int{
	{-1, -1},
	{-1,  0},
	{-1, +1},
	{ 0, -1},
	{ 0, +1},
	{+1, -1},
	{+1,  0},
	{+1, +1},
}

func createGrid(file *os.File) []string {
	var grid  []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	return grid
}

func printGrid(grid []string) {
	for i := range grid {
		log.Println(grid[i])
	}
}

func canClean(grid []string, x int, y int) bool {
	count := 0
	for i := range(d) {
		dx := x+d[i][0]
		dy := y+d[i][1]
		if dx < 0 || dy < 0 || dx >= len(grid[i]) || dy >= len(grid) {
			continue
		}
		if grid[dx][dy] == '@' {
			count++
		}
	}

	return count < 4
}

func solvePart1(file *os.File) {
	grid := createGrid(file)
	printGrid(grid)
	count := 0
	for i := range len(grid) {
		line := grid[i]
		for j := range len(line) {
			if grid[i][j] != '@' {
				continue
			}
			if canClean(grid, i, j) {
				log.Println(grid[i])
				log.Println(i, j)
				count++
			}
		}
	}
	fmt.Println("Solution of part 1:", count)
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

	solvePart1(file)

	file.Close()
}
