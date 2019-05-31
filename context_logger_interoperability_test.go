package log_test

import (
	"context"
	"testing"

	log "github.com/cabify/go-logging"
	"github.com/stretchr/testify/assert"
)

func TestBaggageCanBeReadFromOtherPackages(t *testing.T) {
	ctx := context.Background()
	ctx = log.WithBaggageValue(ctx, "key", "value")

	// notice that we're reading using our own key here
	baggage := ctx.Value("logctx-data-map-string-interface")

	assert.Equal(t, map[string]interface{}{"key": "value"}, baggage)
}
