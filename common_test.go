package gonduit

import (
	"time"

	"github.com/gthvn1/gonduit/util"
)

// t is a wrapper function being used in tests to quickly convert timestamp
// integer into the util.UnixTimestamp object which is supposed to be used in
// response structures.
func timestamp(timestamp int64) util.UnixTimestamp {
	return util.UnixTimestamp(time.Unix(timestamp, 0))
}
