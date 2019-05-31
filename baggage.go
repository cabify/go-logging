package log

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
)

// BaggageContextKey is the key to be used for the baggage map[string]interface{} in
// context.*Value.
// It's intentionally a public string type, so the deprecation of this library is easier, since a public
// string type can be defined in another package
var BaggageContextKey interface{} = "logctx-data-map-string-interface"

const loggerRequestIDLength = 7

// Baggage is a handy helper to extract the baggage from the context, probably to be used with another logger
func Baggage(ctx context.Context) map[string]interface{} {
	if baggage, ok := ctx.Value(BaggageContextKey).(map[string]interface{}); ok {
		return baggage
	}

	return map[string]interface{}{}
}

// WithBaggageValue returns a context with the added key value pair in the baggage store.
func WithBaggageValue(ctx context.Context, key, value string) context.Context {
	oldBaggage, ok := ctx.Value(BaggageContextKey).(map[string]interface{})

	if !ok {
		return context.WithValue(ctx, BaggageContextKey, map[string]interface{}{key: value})
	}

	newBaggage := make(map[string]interface{}, len(oldBaggage)+1)
	for oldKey, oldValue := range oldBaggage {
		newBaggage[oldKey] = oldValue
	}
	newBaggage[key] = value

	return context.WithValue(ctx, BaggageContextKey, newBaggage)
}

// WithBaggageValues returns a context with all key value pairs added to the baggage store.
func WithBaggageValues(ctx context.Context, keyValue map[string]string) context.Context {
	oldBaggage, ok := ctx.Value(BaggageContextKey).(map[string]interface{})
	if !ok {
		return context.WithValue(ctx, BaggageContextKey, toMapStringInterface(keyValue))
	}

	newBaggage := make(map[string]interface{}, len(oldBaggage)+len(keyValue))
	for oldKey, oldValue := range oldBaggage {
		newBaggage[oldKey] = oldValue
	}

	for newKey, newValue := range keyValue {
		newBaggage[newKey] = newValue
	}

	return context.WithValue(ctx, BaggageContextKey, newBaggage)
}

// NewContextWithBaggageFrom returns a new context with baggage values obtained from another context
func NewContextWithBaggageFrom(ctx context.Context) context.Context {
	oldBaggage, ok := ctx.Value(BaggageContextKey).(map[string]interface{})
	if !ok {
		return context.Background()
	}
	return context.WithValue(context.Background(), BaggageContextKey, oldBaggage)
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

func toMapStringInterface(original map[string]string) map[string]interface{} {
	m := make(map[string]interface{}, len(original))
	for k, v := range original {
		m[k] = v
	}
	return m
}
