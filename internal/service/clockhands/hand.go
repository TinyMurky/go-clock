package clockhands

import (
	"math"
	"time"
)

// hand use second as base
type hand struct {
	length float32
	// how much baseTime will the hand travels one round of clock
	timePerRound time.Duration
}

func (h *hand) handInRadius(now time.Time) float32 {
	duration := h.durationSinceStartOfDay(now)
	seconds := math.Mod(duration.Seconds(), h.timePerRound.Seconds())
	radius := (2 * math.Pi * time.Second.Seconds() * seconds) / h.timePerRound.Seconds()

	return float32(radius)
}

func (h *hand) durationSinceStartOfDay(now time.Time) time.Duration {
	beginOfDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	return now.Sub(beginOfDate)
}
