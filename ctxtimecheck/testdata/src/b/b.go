package b

import (
	"time"
	stdtime "time"
)

func ng() {
	time.Date(0, 0, 0, 0, 0, 0, 0, time.Local) // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`
	date := time.Date
	date(0, 0, 0, 0, 0, 0, 0, time.Local) // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`
	func() {
		time.Date(0, 0, 0, 0, 0, 0, 0, time.Local) // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`
		date := time.Date
		date(0, 0, 0, 0, 0, 0, 0, time.Local) // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`
	}()
	stdtime.Date(0, 0, 0, 0, 0, 0, 0, stdtime.Local) // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`

	_ = []struct {
		Time     time.Time
		DateFunc func() time.Time
	}{
		{
			Time: time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`
			DateFunc: func() time.Time {
				return time.Date(0, 0, 0, 0, 0, 0, 0, time.Local) // want `do not use time\.Date, use ctxtime\.Now and its receiver methods to calculate date`
			},
		},
	}
}
