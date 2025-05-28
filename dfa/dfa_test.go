package dfa_test

import (
	"ivanjabrony/fmse/dfa"
	"testing"
)

var tests_dfa = []struct {
	input string
	want  bool
}{
	{"", true},       // 0 единиц (чётное)
	{"0", true},      // 0 единиц
	{"1", false},     // 1 единица (нечётное)
	{"11", true},     // 2 единицы
	{"0110", true},   // 2 единицы
	{"0111", false},  // 3 единицы
	{"111100", true}, // 4 единицы
}

// TestOddOneNumber проверяет количество единиц в слове на четность
func Test_OnesNumber(t *testing.T) {
	dfa := dfa.DFA{
		Alphabet: map[rune]struct{}{
			'0': {},
			'1': {},
		},
		Transitions: map[string]map[rune]string{
			"even": {
				'0': "even", // '0' не меняет чётность
				'1': "odd",  // '1' меняет чётность на нечётную
			},
			"odd": {
				'0': "odd",  // '0' не меняет чётность
				'1': "even", // '1' меняет нечётную на чётную
			},
		},
		Start:  "even",
		Accept: map[string]struct{}{"even": struct{}{}},
	}
	for _, tt := range tests_dfa {
		t.Run(tt.input, func(t *testing.T) {
			ok := dfa.Accepts(tt.input)
			if ok != tt.want {
				t.Errorf("got %v, want %v", ok, tt.want)
			}
		})
	}
}
