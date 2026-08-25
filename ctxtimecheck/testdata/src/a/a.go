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

var pkgFn = func() {
	time.Now() // want `do not use time\.Now, use ctxtime\.Now`
}

var pkgNow = time.Now() // want `do not use time\.Now, use ctxtime\.Now`

var pkgTable = []struct {
	fn func()
}{
	{
		fn: func() {
			time.Now() // want `do not use time\.Now, use ctxtime\.Now`
		},
	},
}
