/*
Package logtest is intended to be used by tests for checking that baggage values have been correctly set
*/
package logtest

import (
	"context"

	"github.com/cabify/go-logging/internal"
)

// HasBaggageValue returns whether the context contains a baggage value.
func HasBaggageValue(ctx context.Context, key, value string) bool {
	baggage, ok := ctx.Value(internal.BaggageContextKey).(internal.Baggage)
	if !ok {
		return false
	}
	return baggage[key] == value
}
