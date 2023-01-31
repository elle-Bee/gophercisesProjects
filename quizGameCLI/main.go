package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// CSV file parsing
	csvFilename := flag.String("csv", "problems.csv", "CSV file in format of 'question,answer' to be importeed")
	timeLimit := flag.Int("timer", 20, "The time limit for the quiz game in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	checkErrNil(err)

	r := csv.NewReader(file)
	listProblems, err := r.ReadAll()
	checkErrNil(err)

	//Making the order of questions random
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(listProblems), func(i, j int) {
		listProblems[i], listProblems[j] = listProblems[j], listProblems[i]
	})

	//Questionaire
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var correct int

quizLoop:
	for i, p := range parseProblems(listProblems) {
		fmt.Printf("Q%d: %s = ", i+1, p.q)
		ansChannel := make(chan string)

		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansChannel <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break quizLoop
		case ans := <-ansChannel:
			if ans == p.a {
				correct++
			}
		}
	}

	fmt.Printf("Congratulations! You've scored %d/%d", correct, len(listProblems))
	defer file.Close()

}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}

type problem struct {
	q string
	a string
}

func parseProblems(listProblems [][]string) []problem {
	returnedProblems := make([]problem, len(listProblems))
	for i, line := range listProblems {
		returnedProblems[i] = problem{
			// Validates the CSV file so that no spaces are counted
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return returnedProblems
}
