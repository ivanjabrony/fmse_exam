package nfa_test

import (
	"ivanjabrony/fmse/nfa"
	"testing"
)

var tests_nfa = []struct {
	name  string
	input string
	want  bool
}{
	{"Empty string", "", false},
	{"Exactly 01", "01", true},
	{"Ends with 01 after many 0s and 1s", "000101101", true},
	{"Ends with 00", "100", false},
	{"Ends with 11", "011", false},
	{"Single 0", "0", false},
	{"Single 1", "1", false},
	{"Long valid string", "1110100010101", true},
	{"Invalid symbol", "012", false},
	{"Only 0s", "0000", false},
	{"Only 1s", "1111", false},
}

// TestOddOneNumber проверяет количество единиц в слове на четность
func Test_EndsWith01(t *testing.T) {
	nfa := &nfa.NFA{
		Alphabet: map[rune]struct{}{
			'0': {},
			'1': {},
		},
		Transitions: map[string]map[rune][]string{
			"q0": {
				'0': {"q0", "q1"}, // Из q0 по '0' можем остаться в q0 или перейти в q1
				'1': {"q0"},       // Из q0 по '1' только в q0
			},
			"q1": {
				'1': {"q2"}, // Из q1 по '1' переходим в принимающее q2
			},
		},
		Start: "q0",
		Accept: map[string]struct{}{
			"q2": {}, // Принимающее состояние
		},
	}
	for _, tt := range tests_nfa {
		t.Run(tt.input, func(t *testing.T) {
			ok := nfa.Accepts(tt.input)
			if ok != tt.want {
				t.Errorf("got %v, want %v", ok, tt.want)
			}
		})
	}
}
