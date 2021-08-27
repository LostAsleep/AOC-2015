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
	// read complete file into memory (newer go just use os.ReadFile())
	dat, err := os.ReadFile(fname)
	checkErr(err)
	return string(dat)
}

const up string = "("
const down string = ")"

func day01() {
	dat := getInput("input.txt")
	floorNr := 0
	for _, direction := range dat {
		direction := string(direction)
		if direction == up {
			floorNr = floorNr + 1
		}
		if direction == down {
			floorNr = floorNr - 1
		}
	}
	fmt.Println("Part 1 - final floor Nr.: ", floorNr)
}

func day02() {
	dat := getInput("input.txt")
	floorNr := 0
	for c, direction := range dat {
		direction := string(direction)
		if direction == up {
			floorNr = floorNr + 1
		}
		if direction == down {
			floorNr = floorNr - 1
		}
		if floorNr == -1 {
			fmt.Println("Part 2 - First time at -1: ", c+1)
			return
		}
	}
}

func main() {
	start := time.Now()

	day01()
	day02()

	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Program took %.3f ms to finish.\n", (float64(elapsed) / 1000))
}
