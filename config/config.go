package config

import "runtime"

var (
	Port = ":8085"
)

var JsonDefault = `{}`

var (
	NumCounters   int64 = 1e7     // Num keys to track frequency of (30M).
	MaxCost       int64 = 1 << 30 // Maximum cost of cache (1GB).
	BufferItems   int64 = 64      // Number of keys per Get buffer.
	NumCPU        int   = runtime.NumCPU()
	TimeOutSearch int   = 15  // secouds
	TTlCache      int   = 600 // secouds => 10 min
)
