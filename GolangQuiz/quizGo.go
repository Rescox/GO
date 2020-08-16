package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file := flag.String("File", "fileQuestions.csv", "Question csv file") //Path csv
	readFile(*file)
}

func readFile(sFile string) {
	secondsCountdown := flag.Int("Seconds", 30, "Time for completing the quiz")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	f, err := os.Open(sFile)
	if err != nil {
		fmt.Printf("Error opening csv file")
	}
	defer f.Close()
	aQuestions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Printf("Error reading csv file")
	}
	iCorrectAnswers := 0
	var enterBegin string
	fmt.Println("Hello and welcome to my golang quiz, type y and Enter to begin or n and Enter to exit...")
	fmt.Scan(&enterBegin)
	if enterBegin == "y" {
		ch := make(chan int)
		go func() {
			var answer string
			fmt.Printf("Goodluck! You have: %v seconds \n", *secondsCountdown)
			fmt.Println("-----------------------")
			for i := 0; i < len(aQuestions); i++ {
				fmt.Println("Question number: ", i)
				fmt.Println(aQuestions[i][0])
				fmt.Scan(&answer)
				if answer == aQuestions[i][1] {
					iCorrectAnswers = iCorrectAnswers + 1
				}
				fmt.Print("\033[2J") //Clearing terminal
			}
			ch <- 1
		}()

		select {
		case <-ch:
			fmt.Printf("The number of correct answers are: %v", iCorrectAnswers)
		case <-time.After(time.Second * time.Duration(*secondsCountdown)):
			fmt.Printf("The number of correct answers are: %v", iCorrectAnswers)
		}
	} else {
		fmt.Println("Goodbye!")
	}
}
