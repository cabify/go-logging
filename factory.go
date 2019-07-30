package log

import (
	"context"
)

var DefaultFactory Factory = NewFactory()

type LoggerFactory interface {
	For(ctx context.Context) Logger
}

// For provides a logger which is aware of the passed context and will prepend
// the context baggage values, using DefaultLogger as base logger.
func For(ctx context.Context) Logger {
	return DefaultFactory.For(ctx)
}
