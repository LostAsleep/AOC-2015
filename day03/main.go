package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func checkErr(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func getInput(fname string) string {
	// read complete file into memory (newer Go just use os.ReadFile())
	dat, err := os.ReadFile(fname)
	checkErr(err)
	return string(dat)
}

const north string = "^"
const south string = "v"
const east string = ">"
const west string = "<"

func move(x int, y int, dir string) (int, int) {
	switch move := dir; move {
	case north:
		y = y + 1
	case south:
		y = y - 1
	case east:
		x = x + 1
	case west:
		x = x - 1
	}
	return x, y
}

type house struct {
	x int
	y int
}

var origin = house{x: 0, y: 0}

func uniqueHouses(input []house) []house {
	// Return a unique subset of the house slice provided.
	// This works because maps can only hold unique keys.
	u := make([]house, 0, len(input))
	m := make(map[house]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

func day1() {
	dat := getInput("input.txt")
	maxHouses := len(dat)
	// Create slice with maximum needed capacity (+ starting point)
	tour := make([]house, 0, maxHouses+1)
	tour = append(tour, origin)
	x, y := origin.x, origin.y
	for _, dir := range dat {
		x, y = move(x, y, string(dir))
		tour = append(tour, house{x, y})
	}
	uHouses := uniqueHouses(tour)
	fmt.Println("Part 1 - Unique Houses:", len(uHouses))
}

func day2() {
	dat := getInput("input.txt")
	maxHouses := len(dat)
	// Create slice with maximum needed capacity (+ starting point)
	tour1 := make([]house, 0, maxHouses)
	tour2 := make([]house, 0, maxHouses)
	tour1 = append(tour1, origin)
	tour2 = append(tour2, origin)
	x1, y1 := origin.x, origin.y
	x2, y2 := origin.x, origin.y

	for i, dir := range dat {
		if i%2 == 0 {
			x1, y1 = move(x1, y1, string(dir))
			tour1 = append(tour1, house{x1, y1})
		} else if i%2 != 0 {
			x2, y2 = move(x2, y2, string(dir))
			tour2 = append(tour2, house{x2, y2})
		}
	}
	tour := append(tour1, tour2...)
	uHouses := uniqueHouses(tour)
	fmt.Println("Part 2 - Unique Houses:", len(uHouses))
}

func main() {
	start := time.Now()

	day1()
	day2()

	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Program took %.3f ms to finish.\n", (float64(elapsed) / 1000))
}
