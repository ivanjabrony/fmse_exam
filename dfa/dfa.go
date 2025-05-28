package dfa

// DFA is a struct that implements basic DFA word acceptance checking.
//
// Possible states are excluded becasue those states can be derived from other fields
type DFA struct {
	Alphabet    map[rune]struct{}
	Transitions map[string]map[rune]string
	Start       string
	Accept      map[string]struct{}
}

// Accepts checks if the word is accepted by DFA or not
func (d *DFA) Accepts(s string) bool {
	current := d.Start
	for _, c := range s {
		if _, ok := d.Alphabet[c]; !ok {
			return false //not supported symbol occured
		}
		next, exists := d.Transitions[current][c]
		if !exists {
			return false //no such transition exists
		}
		current = next
	}

	_, accepts := d.Accept[current] //check if final state is acceptable
	return accepts
}
