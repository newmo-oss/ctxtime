package ctxtimetest_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/newmo-oss/gotestingmock"
	"github.com/newmo-oss/testid"

	"github.com/newmo-oss/ctxtime"
	"github.com/newmo-oss/ctxtime/ctxtimetest"
)

func TestSetFixedNow(t *testing.T) {
	t.Parallel()

	t.Run("before calling SetFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		now1 := ctxtime.Now(ctx)
		time.Sleep(100 * time.Nanosecond)
		now2 := ctxtime.Now(ctx)
		if now1.IsZero() || now2.IsZero() || now1 == now2 || now1.After(now2) {
			t.Errorf("Now must return current time without calling SetFixedNow: %v %v", now1, now2)
		}
	})

	t.Run("after calling SetFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		now := ctxtime.Now(ctx)
		ctxtimetest.SetFixedNow(t, ctx, now)
		fixed := ctxtime.Now(ctx)
		if fixed != now {
			t.Errorf("ctxtime.Now must return the time which had been set by SetFixedNow: %v %v", fixed, now)
		}
	})

	t.Run("after calling UnsetFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		now := ctxtime.Now(ctx)
		ctxtimetest.SetFixedNow(t, ctx, now)
		ctxtimetest.UnsetFixedNow(t, ctx)
		got := ctxtime.Now(ctx)
		if now == got || now.After(got) {
			t.Errorf("ctxtime.Now must return current time after calling UnsetFixedNow: %v %v", got, now)
		}
	})

	t.Run("different test ID", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())

		now1 := ctxtime.Now(ctx)
		time.Sleep(100 * time.Nanosecond)
		now2 := ctxtime.Now(ctx)

		ctxtimetest.SetFixedNow(t, ctx, now1)
		fixed1 := ctxtime.Now(ctx)

		// test IDを変更
		ctx = testid.WithValue(context.Background(), uuid.New().String())
		got := ctxtime.Now(ctx)

		if got == fixed1 || got == now1 {
			t.Errorf("ctxtime.Now must return different time between diffrent test IDs: %v %v", got, fixed1)
		}

		ctxtimetest.SetFixedNow(t, ctx, now2)
		fixed2 := ctxtime.Now(ctx)
		if fixed2 == fixed1 || fixed2 != now2 {
			t.Errorf("ctxtime.Now must return different time between diffrent test IDs: %v %v", fixed2, fixed1)
		}
	})

	t.Run("unset test ID", func(t *testing.T) {
		t.Parallel()

		got := gotestingmock.Run(func(tb *gotestingmock.TB) {
			ctx := context.Background()
			now := ctxtime.Now(ctx)
			ctxtimetest.SetFixedNow(tb, ctx, now)
		})

		if !(got.Failed && got.Goexit) {
			t.Error("ctxtimetest.SetFixedNow must mark failed and exit goroutine of the test when the test id was not related to the context")
		}
	})

	t.Run("unset test ID and not calling SetFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		now1 := ctxtime.Now(ctx)
		time.Sleep(100 * time.Nanosecond)
		now2 := time.Now().In(time.UTC)
		if now1.IsZero() || now2.Before(now1) {
			t.Error("ctxtime.Now must return the current time in unit tests")
		}
	})

	t.Run("set zero time", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		ctxtimetest.SetFixedNow(t, ctx, time.Time{})
		fixed := ctxtime.Now(ctx)
		if !fixed.IsZero() {
			t.Errorf("SetFixedNow must allow zero time")
		}
	})
}
