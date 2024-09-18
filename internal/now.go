package internal

import (
	"context"
	"time"
)

// Now will be changed by functions of newmotimetest package.
var Now = DefaultNow

// DefaultNow is default behavior of [newmotime.Now].
func DefaultNow(_ context.Context) time.Time {
	return time.Now().In(time.UTC)
}
