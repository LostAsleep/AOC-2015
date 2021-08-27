package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func checkError(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func getInput(fname string) string {
	// read complete file into memory (newer Go just use os.ReadFile())
	dat, err := os.ReadFile(fname)
	checkError(err)
	return string(dat)
}

func convStrToInt(str string) int {
	newInt, err := strconv.Atoi(str)
	checkError(err)
	return newInt
}

func getSurfaceAndExtra(l int, w int, h int) int {
	// box surface area: 2*l*w + 2*w*h + 2*h*l
	// slack: area of the smallest side.
	surfaces := []int{l * w, w * h, h * l}
	sort.Ints(surfaces)
	smallestSurface := surfaces[0]
	area := 0
	for _, surface := range surfaces {
		area = area + surface*2
	}
	area = area + smallestSurface
	return area
}

func calcRibbon(l int, w int, h int) int {
	// length = the smallest perimeter of any one face.
	// Plus bow = equal to the cubic feet of volume of present.
	dims := [3]int{l, w, h}
	sort.Ints(dims[:]) // Can only use slices
	smallest := dims[0]
	second_smallest := dims[1]
	ribbon := (smallest*2 + second_smallest*2)
	totalRibbon := ribbon + (l * w * h)
	return totalRibbon
}

func part1() {
	// Get total square feet of wrapping paper.
	dat := getInput("input.txt")
	data := strings.Split(dat, "\n")
	allPackages := 0
	for _, dimensions := range data {
		dim := strings.Split(string(dimensions), "x")
		// convert the dimenstion to integers
		l := convStrToInt(dim[0])
		w := convStrToInt(dim[1])
		h := convStrToInt(dim[2])
		currentSurface := getSurfaceAndExtra(l, w, h)
		allPackages = allPackages + currentSurface
	}
	fmt.Println("Needed wraping paper:", allPackages, "sqFt")
}

func part2() {
	// Get total length of ribbon
	dat := getInput("input.txt")
	data := strings.Split(dat, "\n")
	allPackages := 0
	for _, dimensions := range data {
		dim := strings.Split(string(dimensions), "x")
		// convert the dimenstion to integers
		l := convStrToInt(dim[0])
		w := convStrToInt(dim[1])
		h := convStrToInt(dim[2])
		packageRibbon := calcRibbon(l, w, h)
		allPackages = allPackages + packageRibbon
	}
	fmt.Println("Needed ribbon:", allPackages, "Ft")
}

func main() {
	start := time.Now()

	part1()
	part2()

	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Program took %.3f ms to finish.\n", (float64(elapsed) / 1000))
}
