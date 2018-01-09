package sets

import (
	"fmt"
)

// Set is a key:value type
// Set can only have one item
type Set map[string]struct{}

func main() {
	s := make(Set)
	s["addedItem1"] = struct{}{}
	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var val []string
	for k := range s {
		val = append(val, k)
	}
	return val
}
