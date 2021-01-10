package util

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseBasedOnEmptyLine() []string {
	fN := flag.String("file", "input", "File name")
	flag.Parse()

	file, err := os.Open(*fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := make([]string, 0)

	scanner := bufio.NewScanner(file)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 1, data[0:i], nil
			}
		}

		if atEOF {
			return len(data), data, nil
		}
		// Request more data
		return 0, nil, nil
	}

	scanner.Split(onComma)

	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		report = append(report, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return report
}

func ParseBasedOnEachLine() []string {
	fN := flag.String("file", "input", "File name")
	flag.Parse()

	file, err := os.Open(*fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		report = append(report, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return report
}

func ParseIntobyteArray(fn string) [][]uint8 {
	fN := strings.Trim(fn, " ")
	if strings.EqualFold(fN, "") {
		fN = "input"
	}
	file, err := os.Open(fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := [][]uint8{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report = append(report, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return report
}
