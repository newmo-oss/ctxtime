package newmotime

import (
	"context"
	"time"

	"github.com/newmo-oss/newmotime/internal"
)

// Now returns the current local time in UTC.
// In unit tests, the return value can be controled by functions of newmotimetest package.
func Now(ctx context.Context) time.Time {
	return internal.Now(ctx)
}
