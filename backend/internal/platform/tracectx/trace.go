// Package tracectx demonstrates how request information travels through context.Context.
package tracectx

import (
	"context"
	"fmt"
)

// key is private, preventing other packages from accidentally colliding with it.
type key struct{}

// WithID returns a new context containing the supplied trace ID.
func WithID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, key{}, traceID)
}

// ID returns the trace ID stored in ctx.
func ID(ctx context.Context) (string, bool) {
	traceID, ok := ctx.Value(key{}).(string)
	return traceID, ok
}

// Log prints which layer is currently handling the request.
func Log(ctx context.Context, layer, message string) {
	traceID, ok := ID(ctx)
	if !ok {
		traceID = "no-trace"
	}

	fmt.Printf("[%s] %-12s %s\n", traceID, layer, message)
}
