package dice

import (
	"math/rand"
	"testing"
)

func TestRollNewRoll(t *testing.T) {
	tests := []struct {
		input string
		roll  string
		isErr bool
	}{
		{"", "", false},
		{"d6", "d6", false},
		{"2d6", "2d6", false},
		{"d6,d10,d20,d6,d8", "2d6,d8,d10,d20", false},
		{"x", "", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			roll, err := NewRoll(test.input)

			if test.isErr && err == nil {
				t.Errorf("NewRoll with string %v should have errored", test.input)
			} else if !test.isErr && err != nil {
				t.Errorf("NewRoll with string %v errored unexpectedly", test.input)
			}

			if roll.String() != test.roll {
				t.Errorf("Unexpected Roll from NewRoll: got %v want %v", roll.String(), test.roll)
			}
		})
	}
}

func TestRollRoll(t *testing.T) {
	rolls := [][]Die{
		{Die(6)},
		{Die(6), Die(6)},
		{Die(6), Die(6), Die(6)},
		{Die(1)},
		{Die(1), Die(1), Die(1)},
		{Die(100)},
		{Die(4), Die(6), Die(8), Die(10), Die(20)},
	}

	// 100 random rolls, each with up to 100 dice with values
	// between 1-100.
	for i := 0; i < 100; i++ {
		var r []Die
		d := rand.Intn(100) + 1
		for j := 0; j < d; j++ {
			r = append(r, Die(rand.Intn(100)+1))
		}
		rolls = append(rolls, r)
	}

	for _, r := range rolls {
		roll := Roll{r}
		result := roll.Roll()
		if result.Value < roll.Min() || result.Value > roll.Max() {
			t.Errorf("Roll total %d is outside expected range (%d-%d)", result.Value, roll.Min(), roll.Max())
		}
	}
}

func TestRollMin(t *testing.T) {
	rolls := [][]Die{
		{Die(6)},
		{Die(6), Die(6)},
		{Die(6), Die(6), Die(6)},
		{Die(1)},
		{Die(1), Die(1), Die(1)},
		{Die(100)},
		{Die(4), Die(6), Die(8), Die(10), Die(20)},
	}

	for _, r := range rolls {
		roll := Roll{r}
		t.Run(roll.String(), func(t *testing.T) {
			if roll.Min() != len(r) {
				t.Errorf("Min returned wrong value: got %v wanted %v", roll.Min(), len(r))
			}
		})
	}
}

func TestRollMax(t *testing.T) {
	rolls := [][]Die{
		{Die(6)},
		{Die(6), Die(6)},
		{Die(6), Die(6), Die(6)},
		{Die(1)},
		{Die(1), Die(1), Die(1)},
		{Die(100)},
		{Die(4), Die(6), Die(8), Die(10), Die(20)},
	}

	for _, r := range rolls {
		roll := Roll{r}
		t.Run(roll.String(), func(t *testing.T) {
			var max int
			for _, d := range r {
				max += int(d)
			}

			if roll.Max() != max {
				t.Errorf("Max returned wrong value: got %v wanted %v", roll.Max(), max)
			}
		})
	}
}

func TestRollString(t *testing.T) {
	tests := []struct {
		roll   Roll
		result string
	}{
		{Roll{}, ""},
		{Roll{[]Die{Die(6)}}, "d6"},
		{Roll{[]Die{Die(6), Die(6)}}, "2d6"},
		{Roll{[]Die{Die(6), Die(6), Die(6)}}, "3d6"},
		{Roll{[]Die{Die(6), Die(20)}}, "d6,d20"},
		{Roll{[]Die{Die(6), Die(6), Die(6), Die(20), Die(20)}}, "3d6,2d20"},
		{Roll{[]Die{Die(6), Die(8), Die(10), Die(20)}}, "d6,d8,d10,d20"},
		{Roll{[]Die{Die(6), Die(8), Die(8), Die(8), Die(10)}}, "d6,3d8,d10"},
	}

	for _, test := range tests {
		s := test.roll.String()
		t.Run(s, func(t *testing.T) {
			if s != test.result {
				t.Errorf("Unexpecte String representation of roll: got %v wanted %v", s, test.result)
			}
		})
	}
}

func TestRollParseDice(t *testing.T) {
	tests := []struct {
		s     string
		dice  []Die
		isErr bool
	}{
		{"", []Die{}, false},
		{"d6", []Die{Die(6)}, false},
		{"2d6", []Die{Die(6), Die(6)}, false},
		{"d6,d6", []Die{Die(6), Die(6)}, false},
		{"d6,d20", []Die{Die(6), Die(20)}, false},
		{"d6,2d6", []Die{Die(6), Die(6), Die(6)}, false},
		{"3d6,2d20", []Die{Die(6), Die(6), Die(6), Die(20), Die(20)}, false},
		{"d6,d8,d10,d20", []Die{Die(6), Die(8), Die(10), Die(20)}, false},
		{"d6,3d8,d10", []Die{Die(6), Die(8), Die(8), Die(8), Die(10)}, false},
		// Error states.
		{"dx", []Die{}, true},
		{"x", []Die{}, true},
		{"d0", []Die{}, true},
		{"0d6", []Die{}, true},
	}

	for _, test := range tests {
		dice, err := ParseDice(test.s)

		t.Run(test.s, func(t *testing.T) {
			if test.isErr && err == nil {
				t.Errorf("Dice string %v should have errored on parse", test.s)
			} else if !test.isErr && err != nil {
				t.Errorf("Dice string %v errored on parse: %v", test.s, err)
			}

			if len(dice) != len(test.dice) {
				t.Errorf("Unexpected number of dice on parse: got %v want %v", len(dice), len(test.dice))
			}

			for i, d := range dice {
				if d != test.dice[i] {
					t.Errorf("Unexpected dice on parse at position %d: got %v want %v", i, d, test.dice[i])
				}
			}
		})
	}
}
