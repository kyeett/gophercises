package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	filenamePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\"")
	limitPtr := flag.Int("limit", 10, "the time limit for the quiz in seconds (default 10)")
	flag.Parse()

	f, err := os.Open(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}

	fr := csv.NewReader(f)

	records, err := fr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	questions := parseRecords(records)

	score := runQuiz(questions, *limitPtr)
	fmt.Printf("You scored %d out of %d\n", score, len(questions))
}

func runQuiz(questions []question, limit int) (score int) {

	correctAnswer := make(chan bool)
	quizComplete := make(chan struct{})

	// Goroutine for user input
	go func() {
		for i, question := range questions {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("Problem #%d: %s = ", i+1, question.q)
			//Get answer
			scanner.Scan()
			answer := scanner.Text()
			if strings.Compare(answer, question.a) == 0 {
				correctAnswer <- true
			}
		}
		close(quizComplete)
	}()

	// Accept answers until timeout
	timeout := time.NewTimer(time.Duration(limit) * time.Second)
	for {
		select {
		case <-quizComplete:
			close(correctAnswer)
			return
		case <-timeout.C:
			fmt.Printf("\nTimeout after %d second(s)\n", limit)
			return
		case <-correctAnswer:
			score++
		}
	}
}

func parseRecords(records [][]string) []question {
	var ret []question
	for _, record := range records {
		ret = append(ret, question{
			q: record[0],
			a: record[1],
		})
	}
	return ret
}

type question struct {
	q string
	a string
}
