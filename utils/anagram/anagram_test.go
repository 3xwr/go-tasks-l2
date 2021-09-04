package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreAnagram(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "колба",
			s2:   "бокал",
			want: true,
		},
		{
			s1:   "колба",
			s2:   "охота",
			want: false,
		},
	}

	for _, tc := range tests {
		got := AreAnagram(tc.s1, tc.s2)
		assert.Equal(t, tc.want, got)
	}
}

func TestRemoveDuplicateStr(t *testing.T) {
	got := []string{
		"Листок",
		"Листок",
		"Листок",
		"тяпка",
		"пятка",
		"пятак",
		"пятак",
		"слиток",
		"столик",
	}

	want := []string{
		"Листок",
		"тяпка",
		"пятка",
		"пятак",
		"слиток",
		"столик",
	}

	gotval := removeDuplicateStr(got)
	assert.Equal(t, want, gotval)
}

func TestMakeMapOfSets(t *testing.T) {
	got := []string{
		"Листок",
		"Листок",
		"Листок",
		"тяпка",
		"пятка",
		"пятак",
		"пятак",
		"слиток",
		"столик",
	}

	want := map[string][]string{
		"листок": {"листок", "слиток", "столик"},
		"тяпка":  {"пятак", "пятка", "тяпка"},
	}

	gotval := makeMapOfSets(got)
	assert.Equal(t, want, gotval)
}
