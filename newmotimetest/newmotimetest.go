package newmotimetest

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/newmo-oss/newmotime/internal"
	"github.com/newmo-oss/testid"
)

var fixedNows sync.Map

func init() {
	if testing.Testing() {
		internal.Now = nowForTest
	}
}

// SetFixedNow fixes the return value of newmotime.Now.
// The fixed current time is set each test id which get from [testid.FromContext].
// If any test id cannot obtain from the context, the test will be fail with t.Fatal.
// The fixed current time will be remove by t.Cleanup.
func SetFixedNow(t testing.TB, ctx context.Context, tm time.Time) {
	t.Helper()

	tid, ok := testid.FromContext(ctx)
	if !ok {
		t.Fatal("failed to get test ID from the context")
	}

	t.Cleanup(func() {
		fixedNows.Delete(tid)
	})

	fixedNows.Store(tid, tm)
}

// UnsetFixedNow removes the fixed current time which was set by [SetFixedNow].
// If any test id cannot obtain from the context, the test will be fail with t.Fatal.
func UnsetFixedNow(t testing.TB, ctx context.Context) {
	t.Helper()

	tid, ok := testid.FromContext(ctx)
	if !ok {
		t.Fatal("failed to get test ID from the context")
	}

	fixedNows.Delete(tid)
}

func loadFixedTime(ctx context.Context) (time.Time, bool) {
	tid, ok := testid.FromContext(ctx)
	if !ok {
		return time.Time{}, false
	}

	v, ok := fixedNows.Load(tid)
	if !ok {
		return time.Time{}, false
	}

	tm, ok := v.(time.Time)
	if !ok || tm.IsZero() {
		return time.Time{}, false
	}

	return tm, true
}

func nowForTest(ctx context.Context) time.Time {
	tm, ok := loadFixedTime(ctx)
	if !ok || tm.IsZero() {
		return internal.DefaultNow(ctx)
	}
	return tm
}
