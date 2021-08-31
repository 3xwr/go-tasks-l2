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

func UtilSort(input []byte) []string {
	output := strings.Split(string(input), "\n")

	if b {
		for i, v := range output {
			tmp := strings.TrimSpace(v)
			output[i] = tmp
		}
	}

	if col > 0 {
		tmp := []string{}
		for _, v := range output {
			tmp = append(tmp, strings.Split(v, separator)[col-1])
		}
		output = tmp
	}

	fmt.Println(output)

	if u {
		uniqueMap := make(map[string]struct{})
		for _, v := range output {
			uniqueMap[v] = struct{}{}
		}
		output = []string{}
		for k := range uniqueMap {
			output = append(output, k)
		}
	}

	if m {
		incOrder := func(i, j int) bool {
			return getMonthNumber(output[i]) < getMonthNumber(output[j])
		}
		decOrder := func(i, j int) bool {
			return getMonthNumber(output[i]) > getMonthNumber(output[j])
		}

		if r {
			sort.SliceStable(output, decOrder)
		} else {
			sort.SliceStable(output, incOrder)
		}
	}

	if n || h {
		if c {
			if sort.StringsAreSorted(output) {
				fmt.Println("Sorted")
				return []string{"Sorted"}
			} else {
				fmt.Println("Not Sorted")
				return []string{"Not sorted"}
			}
		}
		natsort.Sort(output)
		if r {
			for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
				output[i], output[j] = output[j], output[i]
			}
		}

	}

	if !n && !h && !m {
		if c {
			if sort.StringsAreSorted(output) {
				fmt.Println("Sorted")
				return []string{"Sorted"}
			} else {
				fmt.Println("Not Sorted")
				return []string{"Not sorted"}
			}
		}
		if r {
			sort.Sort(sort.Reverse(sort.StringSlice(output)))
		} else {
			sort.Strings(output)
		}

	}

	if col > 0 {
		tmp := make(map[string]string)
		for _, v := range strings.Split(string(input), "\n") {
			line := strings.Split(v, separator)
			tmp[line[col-1]] = v
		}
		for i := range output {
			output[i] = tmp[output[i]]
		}
	}

	fmt.Println(output)
	return output
}
