/*
Package logtest is intended to be used by tests for checking that baggage values have been correctly set
*/
package logtest

import (
	"context"

	log "github.com/cabify/go-logging"
)

// HasBaggageValue returns whether the context contains a baggage value.
func HasBaggageValue(ctx context.Context, key, value string) bool {
	baggage, ok := ctx.Value(log.BaggageContextKey).(map[string]interface{})
	if !ok {
		return false
	}
	return baggage[key] == value
}
