package brawlstars

import "time"

type Ratelimit struct {
	Remaining *int32    // requests left for this time window
	Reset     time.Time // unix timestamp until the ratelimit resets
}
