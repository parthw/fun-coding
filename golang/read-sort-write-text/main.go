package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	fileName := "./something.txt"
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	var allText []string
	for {
		text, err := r.ReadString('\n')
		if err == io.EOF {
			allText = append(allText, text)
			break
		}
		if err != nil {
			panic(err)
		}
		allText = append(allText, text)
	}

	sort.Strings(allText)
	fmt.Println(allText)

	newfile, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(newfile)
	for _, s := range allText {
		w.WriteString(s)
	}
	w.Flush()

}
