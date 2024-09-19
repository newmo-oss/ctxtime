package internal

import (
	"context"
	"time"
)

// Now will be changed by functions of ctxtimetest package.
var Now = DefaultNow

// DefaultNow is default behavior of [ctxtime.Now].
func DefaultNow(_ context.Context) time.Time {
	return time.Now().In(time.UTC)
}
