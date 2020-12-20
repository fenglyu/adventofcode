package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInputFile() []string {
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
			return len(data), data[0:len(data)], nil
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

func main() {

	sum := 0
	sum2 := 0

	report := parseInputFile()
	for _, v := range report {
		_, c := countQuestions(v)
		_, c2 := countPart2(v)
		sum += c
		sum2 += c2
		//fmt.Printf("part 2 %d: %v => %d [%d]\n", i, string(r), c2, sum2)
	}
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)

}

func countQuestions(anwser string) ([]uint8, int) {
	bucket := make([]int, 26)
	questions := make([]uint8, 0)

	for _, v := range []byte(anwser) {
		if v < 'a' || v > 'z' {
			continue
		}
		bucket[v-'a']++
	}

	sum := 0

	for i := 0; i < len(bucket); i++ {
		if bucket[i] != 0 {
			questions = append(questions, uint8(i)+'a')
			sum++
		}
		//sum += bucket[i]
	}

	return questions, sum
}

func countPart2(anwser string) ([]uint8, int) {
	//bucket := make([]int, 26)
	questions := make([]uint8, 0)

	lines := strings.Fields(anwser)
	dedup := make([][26]int, len(lines))

	for j, v := range lines {
		for _, f := range []byte(v) {
			if f < 'a' || f > 'z' {
				continue
			}
			dedup[j][f-'a']++
		}
	}

	sum := 0

	for i := 0; i < 26; i++ {

		count := 0
		j := 0
		for j = 0; j < len(dedup); j++ {
			count += dedup[j][i]
		}

		if j == 1 && count > 0 {
			sum++
			questions = append(questions, uint8(i)+'a')
			//  everyone answered yes means the questions count equals the number of people
		} else if count == len(dedup) {
			sum++
			questions = append(questions, uint8(i)+'a')
		}

	}

	return questions, sum
}
