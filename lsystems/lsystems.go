package lsystems

import "math/rand"

type LSystemRule struct {
	Input            rune
	Output           string
	ActivationChance float64
}

type LSystem struct {
	Rules map[rune]LSystemRule
	Axiom rune
}

func (lsys *LSystem) Generate(
	iterations int,
	seed int64,
) string {
	rng := rand.New(rand.NewSource(seed))

	current := string(lsys.Axiom)
	for i := 0; i < iterations; i++ {
		next := ""
		for _, char := range current {
			if rule, found := lsys.Rules[char]; found {
				if rng.Float64() < rule.ActivationChance {
					next += rule.Output
				} else {
					next += string(char)
				}
			} else {
				next += string(char)
			}
		}
		current = next
	}
	return current
}
