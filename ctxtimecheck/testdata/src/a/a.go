package a

import (
	"time"
	stdtime "time"
)

func ng() {
	time.Now() // want `do not use time\.Now, use ctxtime\.Now`
	now := time.Now
	now() // want `do not use time\.Now, use ctxtime\.Now`
	func() {
		time.Now() // want `do not use time\.Now, use ctxtime\.Now`
	}()
	stdtime.Now() // want `do not use time\.Now, use ctxtime\.Now`
}
