package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/facette/natsort"
)

var (
	col                 int
	n, r, u, m, b, c, h bool
	separator           = " "
)

var monthList = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

func init() {
	flag.IntVar(&col, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")

	flag.BoolVar(&m, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учётом суффиксов")
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error while opening file: ", err)
		os.Exit(1)
	}

	UtilSort(input)
}

func getMonthNumber(month string) int {
	month = strings.ToLower(month)
	for i := range monthList {
		if month == strings.ToLower(monthList[i]) {
			return i + 1
		}
	}
	return -1
}

func ignoreSpaces(input []string) []string {
	//func for flag -b
	for i, v := range input {
		tmp := strings.TrimSpace(v)
		input[i] = tmp
	}
	return input
}

func makeUnique(input []string) []string {
	//func for flag -u
	uniqueMap := make(map[string]struct{})
	for _, v := range input {
		uniqueMap[v] = struct{}{}
	}
	input = []string{}
	for k := range uniqueMap {
		input = append(input, k)
	}
	return input
}

func checkIfSorted(input []string) []string {
	//func for flag -c
	if sort.StringsAreSorted(input) {
		fmt.Println("Sorted")
		return []string{"Sorted"}
	} else {
		fmt.Println("Not Sorted")
		return []string{"Not Sorted"}
	}
}

func sortByMonth(input []string, rev bool) []string {
	//func for flag -m and -m -r
	incOrder := func(i, j int) bool {
		return getMonthNumber(input[i]) < getMonthNumber(input[j])
	}
	decOrder := func(i, j int) bool {
		return getMonthNumber(input[i]) > getMonthNumber(input[j])
	}

	if rev {
		sort.SliceStable(input, decOrder)
	} else {
		sort.SliceStable(input, incOrder)
	}

	return input
}

func normalSort(input []string, rev bool) []string {
	//func for basic sorting purposes
	if rev {
		sort.Sort(sort.Reverse(sort.StringSlice(input)))
	} else {
		sort.Strings(input)
	}
	return input
}

func numSort(input []string, rev bool) []string {
	//func for -h and -n flags
	natsort.Sort(input)
	if rev {
		for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]
		}
	}
	return input
}

func changeColOrder(input []string, reqCol int) []string {
	//first part of code for -k flag
	tmp := []string{}
	for _, v := range input {
		tmp = append(tmp, strings.Split(v, separator)[reqCol-1])
	}
	input = tmp

	return input
}

func returnToInitialColOrder(input []byte, output []string, reqCol int) []string {
	tmp := make(map[string]string)
	for _, v := range strings.Split(string(input), "\n") {
		line := strings.Split(v, separator)
		tmp[line[col-1]] = v
	}
	for i := range output {
		output[i] = tmp[output[i]]
	}
	return output
}

func UtilSort(input []byte) []string {
	output := strings.Split(string(input), "\n")

	for _, v := range output {
		fmt.Println(v)
	}

	if b {
		output = ignoreSpaces(output)
	}

	if col > 0 {
		output = changeColOrder(output, col)
		fmt.Println(output)
	}

	if u {
		output = makeUnique(output)
	}

	if m {
		output = sortByMonth(output, r)
	}

	if n || h {
		if c {
			return checkIfSorted(output)
		}
		output = numSort(output, r)
	}

	if !n && !h && !m {
		if c {
			return checkIfSorted(output)
		}
		output = normalSort(output, r)
	}

	if col > 0 {
		output = returnToInitialColOrder(input, output, col)
	}

	fmt.Println(output)
	return output
}
