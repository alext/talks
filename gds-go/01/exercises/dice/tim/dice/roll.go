package dice

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Roll ...
type Roll struct {
	Dice []Die
}

// NewRoll ...
func NewRoll(s string) (Roll, error) {
	dice, err := ParseDice(s)
	if err != nil {
		return Roll{}, err
	}
	return Roll{dice}, nil
}

// RollResult ...
type RollResult struct {
	Value int         `json:"roll"`
	Dice  []DieResult `json:"dice"`
}

// Roller ...
type Roller interface {
	Roll() RollResult
}

// Roll ...
func (r Roll) Roll() RollResult {
	res := RollResult{0, []DieResult{}}

	for _, die := range r.Dice {
		dr := die.Roll()
		res.Value += dr.Value
		res.Dice = append(res.Dice, dr)
	}

	return res
}

// Min ...
func (r Roll) Min() int {
	return len(r.Dice)
}

// Max ...
func (r Roll) Max() int {
	var max int
	for _, d := range r.Dice {
		max += int(d)
	}
	return max
}

// String ...
func (r Roll) String() string {
	dice := make(map[Die]int)
	var dieValues []int

	for _, d := range r.Dice {
		if _, ok := dice[d]; !ok {
			dieValues = append(dieValues, int(d))
		}
		dice[d]++
	}

	// Build a sorted slice of die values so we can output the roll in a
	// predictable order (lowest value to highest).
	keys := make([]int, 0, len(dice))
	for d := range dice {
		keys = append(keys, int(d))
	}
	sort.Ints(keys)

	var s []string
	for _, k := range keys {
		count := dice[Die(k)]
		if count == 1 {
			s = append(s, fmt.Sprintf("d%d", k))
		} else {
			s = append(s, fmt.Sprintf("%dd%d", count, k))
		}
	}

	return strings.Join(s, ",")
}

// ParseDice takes a comma-delimited string of dice definitions and returns
// a representitive slice of Die values.
//
// Example: 3d6 / d6,d6,d6 / d3,d4,d6,5d8
func ParseDice(s string) ([]Die, error) {
	var dice []Die

	if s == "" {
		return dice, nil
	}

	for _, d := range strings.Split(s, ",") {
		d = strings.TrimSpace(d)
		p := strings.Split(d, "d")
		if len(p) != 2 {
			return []Die{}, fmt.Errorf("%s is not a valid dice definition", d)
		}

		value, err := strconv.Atoi(p[1])
		if err != nil || value < 1 {
			return []Die{}, fmt.Errorf("%s is not a valid dice definition (die value is invalid)", d)
		}

		count := 1
		if len(p[0]) > 0 {
			count, err = strconv.Atoi(p[0])
			if err != nil || count < 1 {
				return []Die{}, fmt.Errorf("%s is not a valid dice definition (die count is invalid)", d)
			}
		}

		for i := 0; i < count; i++ {
			dice = append(dice, Die(value))
		}
	}

	return dice, nil
}
