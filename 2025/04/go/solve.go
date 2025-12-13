package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"container/list"
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

func createGrid(file *os.File) [][]uint8 {
	var grid [][]uint8
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []uint8(line))
	}
	return grid
}

func printGrid(grid [][]uint8) {
	for i := range grid {
		log.Println(string(grid[i]))
	}
}

func canClean(grid [][]uint8, x int, y int) bool {
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

func solvePart1(grid [][]uint8) {
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

func solvePart2(grid [][]uint8) {
	q := list.New()
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == '@' {
				q.PushBack([2]int{i, j})
			}
		}
	}

	count := 0
	for q.Len() > 0 {
		coord, _ := q.Remove(q.Front()).([2]int)
		i := coord[0]
		j := coord[1]
		if grid[i][j] == '@' && canClean(grid, i, j) {
			count++
			grid[i][j] = '.'

			for di := range(d) {
				dx := i+d[di][0]
				dy := j+d[di][1]
				if dx < 0 || dy < 0 || dx >= len(grid[i]) || dy >= len(grid) {
					continue
				}
				if grid[dx][dy] == '@' {
					q.PushBack([2]int{dx, dy})
				}
			}
		}
	}
	printGrid(grid)
	fmt.Println("Solution of part 2:", count)
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

	grid := createGrid(file)
	solvePart1(grid)
	solvePart2(grid)

	file.Close()
}
