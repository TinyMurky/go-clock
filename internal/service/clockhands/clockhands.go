package clockhands

import (
	"time"

	"github.com/TinyMurky/go-clock/internal/config"
)

// SecondHand will get the Point of the tail of Second hand at current time
func SecondHand(now time.Time) Point {
	return Point{
		X: config.ClockRadius,
		Y: config.ClockRadius - config.SecondHandLength,
	}
}
