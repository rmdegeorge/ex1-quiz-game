package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// create a help flag that explains the command line arguments
	// that our program will accept
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	// open the file and log the error if there is one
	file, err := os.Open(*csvFilename)

	//error handling
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	// io reader
	r := csv.NewReader(file)

	// read entire file into memory. this is ok here as we know our files
	// won't be big enough to fill up memmory. otherwise we would probably
	// only want to read in lines we were going to use when we need them.
	lines, err := r.ReadAll()

	// exit if there is an error reading the lines
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	// parse the lines of the file into an array called problems
	problems := parseLines(lines)

	// initialize a counter for # of correct answers
	correct := 0

	// loop through each problem
	for i, problem := range problems {
		// Print out the problem
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)

		var answer string
		fmt.Scanf("%s\n", &answer) // Scanf with trim whitespaces so may not be appropriat alwasys
		// check if answer is correct
		if answer == problem.a {
			fmt.Println("Correct!")
			correct++
		}
	}
	fmt.Printf("You scored %d/%d.\n", correct, len(problems))
}

// function that parses the lines of the file into our problem type
// and returns an array of questions and answers
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			// TrimSpace ensures that our answers don't have spaces
			// that would make our answers not match
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// data type for our problems. has a question (q) and an answer (a)
type problem struct {
	q string
	a string
}

// function that exits the program with a message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
