package internal

import (
	"sort"
	"strings"
)

// BaggageContextKey is the key to be used for the Baggage map in
// context.*Value.
const BaggageContextKey ContextKey = "baggage"

// ContextKey is the type of BaggageContextKey.
type ContextKey string

// Baggage is the type of the value associated with BaggageContextKey.
type Baggage map[string]string

// String returns baggages contents as a string of key value pairs ordered alphabetically.
func (b Baggage) String() string {
	var kvPairs []string
	for key, value := range b {
		kvPairs = append(kvPairs, key+":"+value)
	}

	sort.Strings(kvPairs)

	return strings.Join(kvPairs, ": ")
}
