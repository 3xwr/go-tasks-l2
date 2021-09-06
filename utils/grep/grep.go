package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	A, B, C       int
	c, i, v, F, n bool
)

func init() {
	flag.IntVar(&A, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "количество строк")
	flag.BoolVar(&i, "i", false, "игнорировать регистр")
	flag.BoolVar(&v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&F, "F", false, "точное совпадение со строкой, не паттерн")
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

func FixedSearch(input []string, searchPattern string) []int {
	//func for the -F flag
	output := []int{}
	for lineNum, line := range input {
		if strings.Contains(line, searchPattern) {
			output = append(output, lineNum)
		}
	}

	return output
}

func RegExpSearch(input []string, searchPattern string) []int {
	//func for normal search without the -F flag
	output := []int{}
	rExp := regexp.MustCompile(searchPattern)
	for lineNum, line := range input {
		if rExp.MatchString(line) {
			output = append(output, lineNum)
		}
	}

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

func removeDuplicateInts(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func AddBefore(lineNums []int, Before int, appendTo []int) []int {
	for _, v := range appendTo {
		for i := 1; i <= Before; i++ {
			lineNums = append(lineNums, v-i)
		}
	}
	lineNums = removeDuplicateInts(lineNums)
	sort.Ints(lineNums)
	return lineNums
}

func AddAfter(lineNums []int, After int, appendTo []int) []int {
	for _, v := range appendTo {
		for i := 1; i <= After; i++ {
			lineNums = append(lineNums, v+i)
		}
	}
	lineNums = removeDuplicateInts(lineNums)
	sort.Ints(lineNums)
	return lineNums
}

func InvertedSearch(inputArr []int, matchArr []int) []int {
	for i1, v1 := range inputArr {
		for _, v2 := range matchArr {
			if v1 == v2 {
				inputArr = append(inputArr[:i1], inputArr[i1+1:]...)
			}
		}
	}
	return inputArr
}

func GetResult(input []string, inputNums []int) []string {
	output := []string{}
	for _, v := range inputNums {
		output = append(output, input[v])
	}
	return output
}

func AddLineNumber(output []string, inputNums []int, matchArr []int) []string {
	for i, v := range inputNums {
		found := false
		for _, matchNum := range matchArr {
			if v == matchNum {
				output[i] = fmt.Sprintf("%d:%s", v+1, output[i])
				found = true
			}
		}
		if !found {
			output[i] = fmt.Sprintf("%d-%s", v+1, output[i])
		}

	}
	return output
}

func GrepFunc(input []byte, searchPattern string) []string {
	inputString := strings.Split(string(input), "\n")

	var inputStringUntouched = make([]string, len(inputString))
	copy(inputStringUntouched, inputString)

	var lineNums []int
	var outputNums []int

	var inputNums []int
	for i := 0; i < len(inputString); i++ {
		inputNums = append(inputNums, i)
	}

	if i {
		inputString, searchPattern = IgnoreCase(inputString, searchPattern)
	}

	if F {
		lineNums = FixedSearch(inputString, searchPattern)
	} else {
		lineNums = RegExpSearch(inputString, searchPattern)
	}

	if c {
		fmt.Println(len(lineNums))
		num := strconv.Itoa(len(lineNums))
		return []string{num}
	}

	outputNums = lineNums
	//fmt.Println("lineNums before A flag", lineNums)

	if v {
		outputNums = InvertedSearch(inputNums, outputNums)
		//fmt.Println("result after invertion", outputNums)
	}

	if C > 0 {
		A = 0
		B = 0
		outputNums = AddAfter(outputNums, C, lineNums)
		outputNums = AddBefore(outputNums, C, lineNums)
		//fmt.Println(outputNums)
	}

	if A > 0 {
		outputNums = AddAfter(outputNums, A, lineNums)
		//fmt.Println(outputNums)
	}

	if B > 0 {
		outputNums = AddBefore(outputNums, B, lineNums)
		//fmt.Println(outputNums)
	}

	output := GetResult(inputStringUntouched, outputNums)

	if n {
		output = AddLineNumber(output, outputNums, lineNums)
	}

	for _, v := range output {
		fmt.Println(v)
	}

	return inputString
}
