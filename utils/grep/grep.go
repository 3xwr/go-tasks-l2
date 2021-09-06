package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	A, B, C       int
	c, i, v, F, n bool
)

func init() {
	flag.IntVar(&A, "A", 0, "печатать +N строк после совпадения")  //+
	flag.IntVar(&B, "B", 0, "печатать +N строк до совпадения")     //+
	flag.IntVar(&C, "C", 0, "печатать ±N строк вокруг совпадения") //+
	flag.BoolVar(&c, "c", false, "количество строк")
	flag.BoolVar(&i, "i", false, "игнорировать регистр") //+
	flag.BoolVar(&v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&F, "F", false, "точное совпадение со строкой, не паттерн") //+
	flag.BoolVar(&n, "n", false, "печатать номер строки")
}

func main() {
	flag.Parse()

	searchPattern := flag.Arg(0)
	if searchPattern == "" {
		os.Exit(2)
	}
	filename := flag.Arg(1)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error while opening file: ", err)
		os.Exit(1)
	}
	GrepFunc(input, searchPattern)
}

func IgnoreCase(input []string, searchPattern string) ([]string, string) {
	//func for the -i flag
	for i, v := range input {
		v = strings.ToLower(v)
		input[i] = v
	}
	searchPattern = strings.ToLower(searchPattern)
	return input, searchPattern
}

func FixedSearch(input []string, searchPattern string, After, Before, Combined int, inv bool) []string {
	//func for the -F flag
	output := []string{}
	for lineNum, line := range input {
		if strings.Contains(line, searchPattern) {
			if Combined > 0 {
				//code for the -C flag
				CBefore, CAfter := Combined, Combined
				if Before > Combined {
					CBefore = Before
				}

				if After > Combined {
					CAfter = After
				}
				//code for getting lines before
				for i := lineNum - CBefore; i <= lineNum; i++ {
					if i < 0 {
						break
					}
					output = append(output, input[i])
				}
				output = append(output, line)
				//code for getting lines afters
				for i := lineNum + 1; i <= lineNum+CAfter; i++ {
					if i == len(input) {
						break
					}
					output = append(output, input[i])
				}
			}
			//code for the -B flag
			if Before > 0 {
				for i := lineNum - Before; i <= lineNum; i++ {
					if i < 0 {
						break
					}
					output = append(output, input[i])
				}
			}
			output = append(output, line)
			//code for the -A flag
			if After > 0 {
				for i := lineNum + 1; i <= lineNum+After; i++ {
					if i == len(input) {
						break
					}
					output = append(output, input[i])
				}
			}

		}
	}
	output = removeDuplicateStr(output)

	return output
}

func RegExpSearch(input []string, searchPattern string, After, Before, Combined int, inv bool) []string {
	//func for normal search without the -F flag
	output := []string{}
	rExp := regexp.MustCompile(searchPattern)
	for lineNum, line := range input {
		if rExp.MatchString(line) {
			if Combined > 0 {
				//code for the -C flag
				CBefore, CAfter := Combined, Combined
				if Before > Combined {
					CBefore = Before
				}

				if After > Combined {
					CAfter = After
				}
				//code for getting lines before
				for i := lineNum - CBefore; i <= lineNum; i++ {
					if i < 0 {
						break
					}
					output = append(output, input[i])
				}
				output = append(output, line)
				//code for getting lines afters
				for i := lineNum + 1; i <= lineNum+CAfter; i++ {
					if i == len(input) {
						break
					}
					output = append(output, input[i])
				}
			}
			//code for the -B flag
			if Before > 0 {
				for i := lineNum - Before; i <= lineNum; i++ {
					if i < 0 {
						break
					}
					output = append(output, input[i])
				}
			}
			output = append(output, line)
			//code for the -A flag
			if After > 0 {
				for i := lineNum + 1; i <= lineNum+After; i++ {
					if i == len(input) {
						break
					}
					output = append(output, input[i])
				}

			}
		}
	}

	output = removeDuplicateStr(output)

	return output
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GrepFunc(input []byte, searchPattern string) []string {
	output := strings.Split(string(input), "\n")

	if i {
		output, searchPattern = IgnoreCase(output, searchPattern)
	}

	if F {
		output = FixedSearch(output, searchPattern, A, B, C, v)
	} else {
		output = RegExpSearch(output, searchPattern, A, B, C, v)
	}

	for _, v := range output {
		fmt.Println(v)
	}
	return output
}
