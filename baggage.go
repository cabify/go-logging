package log

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/cabify/go-logging/internal"
)

const loggerRequestIDLength = 7

// WithBaggageValue returns a context with the added key value pair in the baggage store.
func WithBaggageValue(ctx context.Context, key, value string) context.Context {
	oldBaggage, ok := ctx.Value(internal.BaggageContextKey).(internal.Baggage)

	if !ok {
		return context.WithValue(ctx, internal.BaggageContextKey, internal.Baggage{key: value})
	}

	newBaggage := make(internal.Baggage, len(oldBaggage)+1)
	for oldKey, oldValue := range oldBaggage {
		newBaggage[oldKey] = oldValue
	}
	newBaggage[key] = value

	return context.WithValue(ctx, internal.BaggageContextKey, newBaggage)
}

// WithBaggageValues returns a context with all key value pairs added to the baggage store.
func WithBaggageValues(ctx context.Context, keyValue map[string]string) context.Context {
	oldBaggage, ok := ctx.Value(internal.BaggageContextKey).(internal.Baggage)
	if !ok {
		return context.WithValue(ctx, internal.BaggageContextKey, internal.Baggage(keyValue))
	}

	newBaggage := make(internal.Baggage, len(oldBaggage)+len(keyValue))
	for oldKey, oldValue := range oldBaggage {
		newBaggage[oldKey] = oldValue
	}

	for newKey, newValue := range keyValue {
		newBaggage[newKey] = newValue
	}

	return context.WithValue(ctx, internal.BaggageContextKey, newBaggage)
}

// NewContextWithBaggageFrom returns a new context with baggage values obtained from another context
func NewContextWithBaggageFrom(ctx context.Context) context.Context {
	oldBaggage, ok := ctx.Value(internal.BaggageContextKey).(internal.Baggage)
	if !ok {
		return context.Background()
	}
	return context.WithValue(context.Background(), internal.BaggageContextKey, oldBaggage)
}

// NewContextFromWithValue creates a new context with baggage values from ctx plus the provided one
// it's just a shorthand for the composition of NewContextWithBaggageFrom and WithBaggageValue
func NewContextFromWithValue(ctx context.Context, k, v string) context.Context {
	return NewContextWithBaggageFrom(
		WithBaggageValue(ctx, k, v),
	)
}

// NewID returns a random id to follow the log traces
func NewID() string {
	data := rand.Int63()
	encoded := md5.Sum([]byte(strconv.FormatInt(data, 16)))
	return fmt.Sprintf("%x", encoded)[:loggerRequestIDLength]
}
