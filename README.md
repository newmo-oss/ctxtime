# ctxtime [![Go Reference](https://pkg.go.dev/badge/github.com/newmo-oss/ctxtimeo.svg)](https://pkg.go.dev/github.com/newmo-oss/ctxtime)[![Go Report Card](https://goreportcard.com/badge/github.com/newmo-oss/ctxtime)](https://goreportcard.com/report/github.com/newmo-oss/ctxtime)

ctxtime provides testable time.Now.

## Usage

By default, `ctxtime.Now(ctx)` returns the current time in UTC.
ctxtimetest.SetFixedNow' can be used to set the return value of `ctxtime.Now(ctx)`.
The return value will be assocciated the test id that can be obtained from the context.
However, if `testing.Testing` returns false, `ctxtimetest.SetFixedNow` won't affect `ctxtime.Now`.

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

## Linter

ctxtimecheck is a linter that finds calls of `time.Now` in your code.

You can install ctxtimecheck by go install.

```sh
$ go install github.com/newmo-oss/ctxtime/ctxtimecheck/cmd/ctxtimecheck@latest
```

ctxtimecheck can be run with the go vet command.

```sh
$ go vet -vettool=$(which ctxtimecheck) ./...
```

If you are already using [gostaticanalysis/called], it can be used instead of ctxtimecheck as follows.

```sh
$ go vet -vettool=$(which called) -called.funcs="time.Now" ./...
```

## License
MIT

[gostaticanalysis/called]: https://github.com/gostaticanalysis/called
