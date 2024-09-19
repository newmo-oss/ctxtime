package ctxtime

import (
	"context"
	"time"

	"github.com/newmo-oss/ctxtime/internal"
)

// Now returns the current local time in UTC.
// In unit tests, the return value can be controled by functions of ctxtimetest package.
func Now(ctx context.Context) time.Time {
	return internal.Now(ctx)
}
