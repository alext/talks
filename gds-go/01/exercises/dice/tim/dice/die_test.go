package dice

import (
	"strconv"
	"testing"
)

func TestDieRoll(t *testing.T) {
	// Test a range of common dice.
	for _, sides := range []int{4, 6, 8, 10, 12, 20, 100} {
		t.Run(strconv.Itoa(sides), func(t *testing.T) {
			die := Die(sides)

			n := 1000 * sides // Make sure we get a good distribution of each side.
			rolls := map[int]int{}

			for i := 0; i < n; i++ {
				roll := die.Roll()
				if roll.Value < 1 || roll.Value > sides {
					t.Errorf("Rolling a d%d returned an invalid result: got %d", sides, roll.Value)
				}
				rolls[roll.Value]++
			}

			// Ensure a reasonable spread of rolls: each side should be rolled 1/sides
			// times, plus or minus 20%.
			delta := (n / sides) / 5
			min := (n / sides) - delta
			max := (n / sides) + delta

			for i := 1; i <= sides; i++ {
				if rolls[i] < min || rolls[i] > max {
					t.Errorf("Unexpected number of rolls of %d on a d%d (expected between %d and %d; got %d)", i, sides, min, max, rolls[i])
				}
			}
		})
	}
}

func TestDieString(t *testing.T) {
	tests := []struct {
		sides int
		s     string
	}{
		{4, "d4"},
		{6, "d6"},
		{8, "d8"},
		{10, "d10"},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			ds := Die(test.sides).String()
			if ds != test.s {
				t.Errorf("Unexpected string representation of die: got %s, want %s)", ds, test.s)
			}
		})
	}
}
