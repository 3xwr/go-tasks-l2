package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreCase(t *testing.T) {
	var tests = []struct {
		input             []string
		searchPattern     string
		wantInput         []string
		wantSearchPattern string
	}{
		{
			input: []string{
				"AAAA",
				"BBBBB",
				"CCC",
			},
			searchPattern: "AaAa",
			wantInput: []string{
				"aaaa",
				"bbbbb",
				"ccc",
			},
			wantSearchPattern: "aaaa",
		},
	}

	for _, tc := range tests {
		gotInput, gotPattern := IgnoreCase(tc.input, tc.searchPattern)
		assert.Equal(t, tc.wantInput, gotInput)
		assert.Equal(t, tc.wantSearchPattern, gotPattern)
	}
}

func TestFixedSearch(t *testing.T) {
	var tests = []struct {
		input         []string
		searchPattern string
		want          []int
	}{
		{
			input: []string{
				"AAAA",
				"BBBBB",
				"CCC",
			},
			searchPattern: "AAAA",
			want:          []int{0},
		},
		{
			input: []string{
				"AAAA",
				"BBBBB",
				"CCC",
			},
			searchPattern: "aaaa",
			want:          []int{},
		},
	}

	for _, tc := range tests {
		got := FixedSearch(tc.input, tc.searchPattern)
		assert.Equal(t, tc.want, got)
	}
}

func TestRegExpSearch(t *testing.T) {
	var tests = []struct {
		input         []string
		searchPattern string
		want          []int
	}{
		{
			input: []string{
				"AAAA",
				"BBBBB",
				"CCC",
			},
			searchPattern: "A+",
			want:          []int{0},
		},
		{
			input: []string{
				"AAAA",
				"BBBBB",
				"CCC",
			},
			searchPattern: ".+",
			want:          []int{0, 1, 2},
		},
	}

	for _, tc := range tests {
		got := RegExpSearch(tc.input, tc.searchPattern)
		assert.Equal(t, tc.want, got)
	}
}

func TestAddBefore(t *testing.T) {
	var tests = []struct {
		lineNums []int
		Before   int
		appendTo []int
		want     []int
	}{
		{
			lineNums: []int{
				2, 4, 6,
			},
			Before: 1,
			appendTo: []int{
				2, 4, 6,
			},
			want: []int{
				1, 2, 3, 4, 5, 6,
			},
		},
		{
			lineNums: []int{
				2, 4, 6,
			},
			Before: 2,
			appendTo: []int{
				2, 4, 6,
			},
			want: []int{
				0, 1, 2, 3, 4, 5, 6,
			},
		},
	}

	for _, tc := range tests {
		got := AddBefore(tc.lineNums, tc.Before, tc.appendTo)
		assert.Equal(t, tc.want, got)
	}
}

func TestAddAfter(t *testing.T) {
	var tests = []struct {
		lineNums []int
		Before   int
		appendTo []int
		want     []int
	}{
		{
			lineNums: []int{
				2, 4, 6,
			},
			Before: 1,
			appendTo: []int{
				2, 4, 6,
			},
			want: []int{
				2, 3, 4, 5, 6, 7,
			},
		},
		{
			lineNums: []int{
				2, 4, 6,
			},
			Before: 2,
			appendTo: []int{
				2, 4, 6,
			},
			want: []int{
				2, 3, 4, 5, 6, 7, 8,
			},
		},
	}

	for _, tc := range tests {
		got := AddAfter(tc.lineNums, tc.Before, tc.appendTo)
		assert.Equal(t, tc.want, got)
	}
}

func TestInvertedSearch(t *testing.T) {
	var tests = []struct {
		inputArr []int
		matchArr []int
		want     []int
	}{
		{
			inputArr: []int{
				0, 1, 2, 3, 4, 5, 6,
			},
			matchArr: []int{
				1, 3, 5,
			},
			want: []int{
				0, 2, 4, 6,
			},
		},
		{
			inputArr: []int{
				0, 1, 2, 3, 4, 5, 6,
			},
			matchArr: []int{
				1,
			},
			want: []int{
				0, 2, 3, 4, 5, 6,
			},
		},
	}

	for _, tc := range tests {
		got := InvertedSearch(tc.inputArr, tc.matchArr)
		assert.Equal(t, tc.want, got)
	}
}

func TestGetResult(t *testing.T) {
	var tests = []struct {
		input     []string
		inputNums []int
		want      []string
	}{
		{
			input: []string{
				"a",
				"b",
				"c",
				"d",
			},
			inputNums: []int{
				0, 3,
			},
			want: []string{
				"a", "d",
			},
		},
		{
			input: []string{
				"a",
				"b",
				"c",
				"d",
			},
			inputNums: []int{
				0, 1, 2, 3,
			},
			want: []string{
				"a", "b", "c", "d",
			},
		},
	}

	for _, tc := range tests {
		got := GetResult(tc.input, tc.inputNums)
		assert.Equal(t, tc.want, got)
	}
}

func TestAddLineNumber(t *testing.T) {
	var tests = []struct {
		output    []string
		inputNums []int
		matchArr  []int
		want      []string
	}{
		{
			output: []string{
				"a",
				"b",
			},
			inputNums: []int{
				0, 3,
			},
			matchArr: []int{
				0, 3,
			},
			want: []string{
				"1:a",
				"4:b",
			},
		},
		{
			output: []string{
				"a",
				"c",
				"b",
			},
			inputNums: []int{
				0, 1, 3,
			},
			matchArr: []int{
				0, 3,
			},
			want: []string{
				"1:a",
				"2-c",
				"4:b",
			},
		},
	}

	for _, tc := range tests {
		got := AddLineNumber(tc.output, tc.inputNums, tc.matchArr)
		assert.Equal(t, tc.want, got)
	}
}
