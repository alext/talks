package dice

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Die ...
type Die int

// DieResult ...
type DieResult struct {
	Type  string `json:"type"`
	Value int    `json:"roll"`
}

// Roll ...
func (d Die) Roll() DieResult {
	return DieResult{d.String(), rand.Intn(int(d)) + 1}
}

// Type ...
func (d Die) String() string {
	return fmt.Sprintf("d%d", d)
}
