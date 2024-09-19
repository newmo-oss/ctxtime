# ctxtime [![Go Reference](https://pkg.go.dev/badge/github.com/newmo-oss/ctxtimeo.svg)](https://pkg.go.dev/github.com/newmo-oss/ctxtime)[![Go Report Card](https://goreportcard.com/badge/github.com/newmo-oss/ctxtime)](https://goreportcard.com/report/github.com/newmo-oss/ctxtime)

ctxtime provides testable time.Now.

## Usage

```go
package a_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/newmo-oss/ctxtime"
	"github.com/newmo-oss/ctxtime/ctxtimetest"
	"github.com/newmo-oss/testid"
)

func Test(t *testing.T) {
	tid := uuid.NewString()
	ctx := testid.WithValue(context.Background(), tid)
	now := ctxtime.Now(ctx)
	ctxtimetest.SetFixedNow(t, ctx, now)
	time.Sleep(10 * time.Second)
	now2 := ctxtime.Now(ctx)
	t.Log(now == now2) // true
}
```

## License
MIT
