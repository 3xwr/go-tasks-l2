package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMonthNumber(t *testing.T) {
	var tests = []struct {
		got  string
		want int
	}{
		{
			got:  monthList[0],
			want: 1,
		},
		{
			got:  monthList[1],
			want: 2,
		},
		{
			got:  monthList[2],
			want: 3,
		},
		{
			got:  monthList[3],
			want: 4,
		},
		{
			got:  monthList[4],
			want: 5,
		},
		{
			got:  monthList[5],
			want: 6,
		},
		{
			got:  monthList[6],
			want: 7,
		},
		{
			got:  monthList[7],
			want: 8,
		},
		{
			got:  monthList[8],
			want: 9,
		},
		{
			got:  monthList[9],
			want: 10,
		},
		{
			got:  monthList[10],
			want: 11,
		},
		{
			got:  monthList[11],
			want: 12,
		},
		{
			got:  "adasdasdasd",
			want: -1,
		},
	}

	for _, tc := range tests {
		gotval := getMonthNumber(tc.got)
		assert.Equal(t, tc.want, gotval)
	}
}

func TestIgnoreSpaces(t *testing.T) {
	var tests = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"a  ",
				"b ",
				"c     ",
			},
			want: []string{
				"a",
				"b",
				"c",
			},
		},
		{
			got: []string{
				"1",
				"2",
				"3",
			},
			want: []string{
				"1",
				"2",
				"3",
			},
		},
	}

	for _, tc := range tests {
		gotval := ignoreSpaces(tc.got)
		assert.Equal(t, tc.want, gotval)
	}
}

func TestMakeUnique(t *testing.T) {
	var tests = []struct {
		got  []string
		want int
	}{
		{
			got: []string{
				"1",
				"1",
				"3",
			},
			want: 2,
		},
		{
			got: []string{
				"1",
				"2",
				"3",
			},
			want: 3,
		},
	}

	for _, tc := range tests {
		gotval := len(makeUnique(tc.got))
		assert.Equal(t, tc.want, gotval)
	}
}

func TestCheckIfSorted(t *testing.T) {
	var tests = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"1",
				"2",
				"3",
			},
			want: []string{
				"Sorted",
			},
		},
		{
			got: []string{
				"5",
				"2",
				"3",
			},
			want: []string{
				"Not Sorted",
			},
		},
	}

	for _, tc := range tests {
		gotval := checkIfSorted(tc.got)
		assert.Equal(t, tc.want, gotval)
	}
}

func TestSortByMonth(t *testing.T) {
	var tests = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"May",
				"Feb",
				"Oct",
			},
			want: []string{
				"Feb",
				"May",
				"Oct",
			},
		},
		{
			got: []string{
				"Jan",
				"Feb",
				"Mar",
			},
			want: []string{
				"Jan",
				"Feb",
				"Mar",
			},
		},
		{
			got: []string{
				"3",
				"2",
				"3",
			},
			want: []string{
				"3",
				"2",
				"3",
			},
		},
	}

	var testsReversed = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"May",
				"Feb",
				"Oct",
			},
			want: []string{
				"Oct",
				"May",
				"Feb",
			},
		},
		{
			got: []string{
				"Jan",
				"Feb",
				"Mar",
			},
			want: []string{
				"Mar",
				"Feb",
				"Jan",
			},
		},
		{
			got: []string{
				"3",
				"2",
				"3",
			},
			want: []string{
				"3",
				"2",
				"3",
			},
		},
	}

	for _, tc := range tests {
		gotval := sortByMonth(tc.got, false)
		assert.Equal(t, tc.want, gotval)
	}

	for _, tc := range testsReversed {
		gotvalRev := sortByMonth(tc.got, true)
		assert.Equal(t, tc.want, gotvalRev)
	}
}

func TestNormalSort(t *testing.T) {
	var tests = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"1",
				"2",
				"10",
			},
			want: []string{
				"1",
				"10",
				"2",
			},
		},
		{
			got: []string{
				"Jan",
				"Feb",
				"Mar",
			},
			want: []string{
				"Feb",
				"Jan",
				"Mar",
			},
		},
		{
			got: []string{
				"3",
				"2",
				"3",
			},
			want: []string{
				"2",
				"3",
				"3",
			},
		},
	}

	var testsReversed = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"1",
				"2",
				"10",
			},
			want: []string{
				"2",
				"10",
				"1",
			},
		},
		{
			got: []string{
				"Jan",
				"Feb",
				"Mar",
			},
			want: []string{
				"Mar",
				"Jan",
				"Feb",
			},
		},
		{
			got: []string{
				"3",
				"2",
				"3",
			},
			want: []string{
				"3",
				"3",
				"2",
			},
		},
	}

	for _, tc := range tests {
		gotval := normalSort(tc.got, false)
		assert.Equal(t, tc.want, gotval)
	}

	for _, tc := range testsReversed {
		gotvalRev := normalSort(tc.got, true)
		assert.Equal(t, tc.want, gotvalRev)
	}
}

func TestNumSort(t *testing.T) {
	var tests = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"1",
				"2",
				"10",
			},
			want: []string{
				"1",
				"2",
				"10",
			},
		},
		{
			got: []string{
				"Jan",
				"Feb",
				"Mar",
			},
			want: []string{
				"Feb",
				"Jan",
				"Mar",
			},
		},
		{
			got: []string{
				"3",
				"2",
				"3",
			},
			want: []string{
				"2",
				"3",
				"3",
			},
		},
	}

	var testsReversed = []struct {
		got  []string
		want []string
	}{
		{
			got: []string{
				"1",
				"2",
				"10",
			},
			want: []string{
				"10",
				"2",
				"1",
			},
		},
		{
			got: []string{
				"Jan",
				"Feb",
				"Mar",
			},
			want: []string{
				"Mar",
				"Jan",
				"Feb",
			},
		},
		{
			got: []string{
				"3",
				"2",
				"3",
			},
			want: []string{
				"3",
				"3",
				"2",
			},
		},
	}

	for _, tc := range tests {
		gotval := numSort(tc.got, false)
		assert.Equal(t, tc.want, gotval)
	}

	for _, tc := range testsReversed {
		gotvalRev := numSort(tc.got, true)
		assert.Equal(t, tc.want, gotvalRev)
	}
}

func TestChangeColOrder(t *testing.T) {
	var tests = []struct {
		got  []string
		col  int
		want []string
	}{
		{
			got: []string{
				"1 c",
				"2 b",
				"3 a",
			},
			col: 2,
			want: []string{
				"c",
				"b",
				"a",
			},
		},
		{
			got: []string{
				"c 1",
				"b 2",
				"a 3",
			},
			col: 2,
			want: []string{
				"1",
				"2",
				"3",
			},
		},
		{
			got: []string{
				"1 a",
				"2 b",
				"3 c",
			},
			col: 2,
			want: []string{
				"a",
				"b",
				"c",
			},
		},
	}

	for _, tc := range tests {
		gotval := changeColOrder(tc.got, tc.col)
		assert.Equal(t, tc.want, gotval)
	}
}
