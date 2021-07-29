package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type marks struct {
	total, correct int
}

func startQuiz(problemFile *os.File, m *marks, inputTime int) {
	reader := csv.NewReader(problemFile)
	qaslice, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	m.total, m.correct = len(qaslice), 0
	timeout := time.After(time.Duration(inputTime) * time.Second)
	for i, qa := range qaslice {
		select {
		case <-timeout:
			return
		default:
			fmt.Printf("%v Please answer %v ", i+1, qa[0])
			var intputAnswer string
			fmt.Scanln(&intputAnswer)

			if intputAnswer == qa[1] {
				m.correct++
			}
		}
	}

}

func main() {
	inputFile := flag.String("inputFile", "./problems.csv", "input csv file")
	inputTime := flag.Int("inputTime", 10, "timeout time in seconds")
	flag.Parse()

	problemFile, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalln("failed to open file")
	}
	defer problemFile.Close()

	var m marks
	startQuiz(problemFile, &m, *inputTime)
	fmt.Println(m.correct, m.total)
}
