package log

import (
	"context"

)

// Factory provides context aware loggers.
type Factory struct {
	baseLogger Logger
}

// NewFactory instantiates a factory with the default logger.
func NewFactory() Factory {
	return Factory{
		baseLogger: DefaultLogger,
	}
}

// For provides a logger which is aware of the passed context and will prepend the context baggage values.
func (f Factory) For(ctx context.Context) Logger {
	return newBaggageLogger(ctx, f.baseLogger)
}

// For provides a logger which is aware of the passed context and will prepend
// the context baggage values, using DefaultLogger as base logger.
func For(ctx context.Context) Logger {
	return newBaggageLogger(ctx, DefaultLogger)
}
