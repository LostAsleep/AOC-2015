package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

const key string = "yzbqklnj"

func part1() string {
	// find MD5 hashes which, in hexadecimal, start with at least 00000
	suffix := 0
	var md5Sum string
	for {
		stringToTest := key + fmt.Sprint(suffix)
		bytesToHash := []byte(stringToTest)
		md5Sum = fmt.Sprintf("%x", md5.Sum(bytesToHash))
		//fmt.Println("String:", stringToTest, " |  md5sum:", md5Sum)
		if md5Sum[0:5] == "00000" {
			break
		}
		suffix = suffix + 1
	}
	return fmt.Sprint(suffix)
}

func part2() string {
	// find MD5 hashes which, in hexadecimal, start with at least 000000
	suffix := 0
	var md5Sum string
	for {
		stringToTest := key + fmt.Sprint(suffix)
		bytesToHash := []byte(stringToTest)
		md5Sum = fmt.Sprintf("%x", md5.Sum(bytesToHash))
		//fmt.Println("String:", stringToTest, " |  md5sum:", md5Sum)
		if md5Sum[0:6] == "000000" {
			break
		}
		suffix = suffix + 1
	}
	return fmt.Sprint(suffix)
}

func main() {
	start := time.Now()
	fiveLeadingZeros := part1()
	fmt.Println(fiveLeadingZeros)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Part1 took %.3f ms to finish.\n", (float64(elapsed) / 1000))

	start = time.Now()
	sixLeadingZeros := part2()
	fmt.Println(sixLeadingZeros)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Part2 took %.3f ms to finish.\n", (float64(elapsed) / 1000))
}
