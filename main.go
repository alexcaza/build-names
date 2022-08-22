package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkValid(e error) {
	if e != nil {
		panic(e)
	}
}

func randomFromFile(path string) (string) {
	// Change seed each time it's run
	rand.Seed(time.Now().UnixNano())

	// Open file for reading
	file, err := os.Open(path)
	checkValid(err)
	defer file.Close()

	// Store lines in array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	checkValid(scanner.Err())

	// Return random line
	return lines[rand.Intn(len(lines))]
}

func getAdjective() (string) {
	return randomFromFile("data/adjectives.txt")
}

func getNoun() (string) {
	return randomFromFile("data/nouns.txt")
}

func padLeft(s string, n int) (string) {
	return strings.Repeat("0", n - len(s)) + s
}

func main() {
	// Setup variables
	var output []string
	var dateString string

	randIntPtr := flag.Bool("r", false, "Include a random number in the name")
	intLenPtr := flag.Int("rl", 3, "Length of the random number (max 6; default 3)")
	datePtr := flag.Bool("d", false, "Include today's date (yyyy-mm-dd) in the name")
	numberPtr := flag.Int("n", 1, "How many names to generate (default 1)")
	alliteratePtr := flag.Bool("a", false, "Generate alliterative names (ex: abject-animal)")

	// Set parse flags and fill pointer addresses
	flag.Parse()

	if *datePtr {
		// Get today's date and set it once to save cycles
		dateString = time.Now().Format("2006-01-02")
	}

	if *intLenPtr > 6 {
		// Exit with error code 1 if int length is too long
		fmt.Println("Integer length cannot be greater than 6")
		os.Exit(1)
	}

	// Loop over the count of the number of names to generate
	for i := 0; i < *numberPtr; i++ {
		// Set base output string
		adj := getAdjective()
		noun := getNoun()

		if *alliteratePtr && adj[0] != noun[0] {
			for {
				noun = getNoun()
				if adj[0] == noun[0] {
					break
				}
			}
		}

		outputString := adj + "-" + noun

		// If the random number flag is set, add a random number to the name
		if *randIntPtr {
			pwr := int(math.Pow(10, float64(*intLenPtr))) - 1
			numAsString := strconv.Itoa(rand.Intn(pwr))
			outputString += "-" + padLeft(numAsString, *intLenPtr)
		}

		// If the date flag is set, add the date to the name
		if *datePtr {
			outputString += "-" + dateString
		}

		// Append the name to the output array
		output = append(output, outputString)
	}

	// Print the names to the console
	if len(output) > 0 {
		fmt.Println(strings.Join(output, "\n"))
	} else {
		fmt.Println(strings.Join(output, ""))
	}
}