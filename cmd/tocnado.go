package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Headline struct {
	rawText     string
	text        string
	level       int
	chainedText string
	headlineID  string
	anchorLink  string
}

func newHeadline(rawHeadline string) Headline {
	headlinePattern := `(#+)\s(.*)`
	headlineRegex := regexp.MustCompile(headlinePattern)

	headlineParts := headlineRegex.FindStringSubmatch(rawHeadline)

	headline := Headline{
		rawText: rawHeadline,
		text:    string(headlineParts[2]),
		level:   len(headlineParts[1]),
	}

	headline.chainedText = strings.ToLower(strings.Replace(headline.text, " ", "-", -1))
	headline.headlineID = "#" + headline.chainedText
	headline.anchorLink = fmt.Sprintf("[%s](%s)", headline.text, headline.headlineID)

	return headline
}

type TableOfContent struct {
	rawHeadlines []Headline
}

// func newTableOfContent(headlines []Headline) TableOfContent {
//
// 	tableOfContent := TableOfContent{
// 		rawHeadlines: headlines,
// 	}
// }

func createBulletPoint(level int) string {

	bulletPoint := ""

	for i := 1; i < level; i++ {
		bulletPoint += "\t"
	}

	bulletPoint += "-"

	return bulletPoint
}

func printTableOfContent(headlines []Headline) {
	for _, headline := range headlines {
		if headline.level == 1 {
			continue
		}
		offset := headline.level - 1
		fmt.Println(createBulletPoint(offset), headline.anchorLink)
	}
}

func getFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getHeadlines(lines []string) []Headline {
	headlinePattern := `(#+)\s(.*)`
	headlineRegex := regexp.MustCompile(headlinePattern)

	var headlines []Headline
	for _, line := range lines {
		if headlineRegex.MatchString(line) {
			headline := newHeadline(line)
			headlines = append(headlines, headline)
		}
	}
	return headlines
}

func validateInput(arg string) {

	if !strings.HasSuffix(arg, ".md") {
		fmt.Println("Please provide a markdown file")
		os.Exit(1)
	}

	if _, err := os.Stat(arg); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	filePath := os.Args[1]

	validateInput(filePath)

	lines, err := getFileLines(filePath)
	if err != nil {
		fmt.Println(err)
	}

	headlines := getHeadlines(lines)

	printTableOfContent(headlines)
}
