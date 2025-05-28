package nfa

// NFA is a struct that implements basic NFA word acceptance checking.
//
// Possible states are excluded becasue those states can be derived from other fields
type NFA struct {
	Alphabet    map[rune]struct{}
	Transitions map[string]map[rune][]string // now it is possible to have multiple transitions from the same state
	Start       string
	Accept      map[string]struct{}
}

// Accepts checks if the word is accepted by NFA or not
func (n *NFA) Accepts(s string) bool {
	currentStates := map[string]struct{}{n.Start: {}}

	for _, c := range s {
		if _, ok := n.Alphabet[c]; !ok {
			return false
		}

		nextStates := make(map[string]struct{})
		for state := range currentStates {
			if transitions, ok := n.Transitions[state][c]; ok {
				for _, next := range transitions {
					nextStates[next] = struct{}{}
				}
			}
		}
		currentStates = nextStates
		if len(currentStates) == 0 {
			return false
		}
	}

	for state := range currentStates {
		if _, ok := n.Accept[state]; ok {
			return true
		}
	}
	return false
}
