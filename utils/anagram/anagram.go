package main

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func main() {
	var input = []string{
		"листок",
		"тяпка",
		"пятка",
		"пятак",
		"слиток",
		"столик",
	}
	fmt.Println(makeMapOfSets(input))
}

type Letter []rune

func (r Letter) Len() int           { return len(r) }
func (r Letter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Letter) Less(i, j int) bool { return r[i] < r[j] }

func AreAnagram(s1 string, s2 string) bool {
	//check if s1 is an anagram of s2
	var r1, r2 Letter
	for _, letter := range s1 {
		r1 = append(r1, letter)
	}
	for _, letter := range s2 {
		r2 = append(r2, letter)
	}

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
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

func makeMapOfSets(input []string) map[string][]string {

	//return set
	tmpSet := make(map[string][]string)

	for _, word := range input {
		isAnagram := false
		word = strings.ToLower(word)
		if len(tmpSet) == 0 {
			//set the first word as key
			tmpSet[word] = []string{word}
			continue
		}
		for k := range tmpSet {
			//check if current word is an anagram to any word in the map
			if AreAnagram(k, word) {
				arr := tmpSet[k]
				arr = append(arr, word)
				tmpSet[k] = arr
				isAnagram = true
			}
		}

		if !isAnagram {
			tmpSet[word] = []string{word}
		}

	}

	for k, v := range tmpSet {
		//remove results that only show up once
		if len(v) < 2 {
			delete(tmpSet, k)
		}
		//remove duplicate strings
		tmpSet[k] = removeDuplicateStr(v)
		//collate is used for alphabetical sorting of strings in russian
		c := collate.New(language.Russian)
		c.SortStrings(tmpSet[k])
	}

	return tmpSet
}
