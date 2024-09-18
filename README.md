# newmotime [![Go Reference](https://pkg.go.dev/badge/github.com/newmo-oss/newmotimeo.svg)](https://pkg.go.dev/github.com/newmo-oss/newmotime)[![Go Report Card](https://goreportcard.com/badge/github.com/newmo-oss/newmotime)](https://goreportcard.com/report/github.com/newmo-oss/newmotime)

newmotime privides testable time.Now

## Usage

```go
package a_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/newmo-oss/newmotime"
	"github.com/newmo-oss/newmotime/newmotimetest"
	"github.com/newmo-oss/testid"
)

func Test(t *testing.T) {
	tid := uuid.NewString()
	ctx := testid.WithValue(context.Background(), tid)
	now := newmotime.Now(ctx)
	newmotimetest.SetFixedNow(t, ctx, now)
	time.Sleep(10 * time.Second)
	now2 := newmotime.Now(ctx)
	t.Log(now == now2) // true
}
```

## License
MIT
