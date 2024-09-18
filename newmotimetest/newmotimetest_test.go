package newmotimetest_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/newmo-oss/testid"
	"github.com/newmo-oss/newmotime"
	"github.com/newmo-oss/newmotime/newmotimetest"
)

func TestWithFixedNow(t *testing.T) {
	t.Parallel()

	t.Run("before calling WithFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		now1 := newmotime.Now(ctx)
		time.Sleep(1 * time.Nanosecond)
		now2 := newmotime.Now(ctx)
		if now1 == now2 || now1.After(now2) {
			t.Errorf("Now must return current time without calling SetFixedNow: %v %v", now1, now2)
		}
	})

	t.Run("after calling WithFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		now := newmotime.Now(ctx)
		newmotimetest.SetFixedNow(t, ctx, now)
		fixed := newmotime.Now(ctx)
		if fixed != now {
			t.Errorf("newmotime.Now must return the time which had been set by SetFixedNow: %v %v", fixed, now)
		}
	})

	t.Run("after calling WithoutFixedNow", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())
		now := newmotime.Now(ctx)
		newmotimetest.SetFixedNow(t, ctx, now)
		newmotimetest.UnsetFixedNow(t, ctx)
		got := newmotime.Now(ctx)
		if now == got || now.After(got) {
			t.Errorf("newmotime.Now must return current time after calling WithoutFixedNow: %v %v", got, now)
		}
	})

	t.Run("different test ID", func(t *testing.T) {
		t.Parallel()

		ctx := testid.WithValue(context.Background(), uuid.New().String())

		now1 := newmotime.Now(ctx)
		time.Sleep(1 * time.Nanosecond)
		now2 := newmotime.Now(ctx)

		newmotimetest.SetFixedNow(t, ctx, now1)
		fixed1 := newmotime.Now(ctx)

		// test IDを変更
		ctx = testid.WithValue(context.Background(), uuid.New().String())
		got := newmotime.Now(ctx)

		if got == fixed1 || got == now1 {
			t.Errorf("newmotime.Now must return different time between diffrent test IDs: %v %v", got, fixed1)
		}

		newmotimetest.SetFixedNow(t, ctx, now2)
		fixed2 := newmotime.Now(ctx)
		if fixed2 == fixed1 || fixed2 != now2 {
			t.Errorf("newmotime.Now must return different time between diffrent test IDs: %v %v", fixed2, fixed1)
		}
	})

	t.Run("unset test ID", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		fakeT := &testingT{T: t}
		now := newmotime.Now(ctx)
		newmotimetest.SetFixedNow(fakeT, ctx, now)
		if !fakeT.callFailNow {
			t.Error("newmotimetest.SetFixedNow must call t.Fatal/t.Fatalf/t.FailNow when test id was not related to the context")
		}
	})
}

type testingT struct {
	*testing.T
	callFailNow bool
}

func (t *testingT) FailNow() {
	t.callFailNow = true
}

func (t *testingT) Fatal(args ...any) {
	t.callFailNow = true
}

func (t *testingT) Fatalf(format string, args ...any) {
	t.callFailNow = true
}
