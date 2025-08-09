package clockhands_test

import (
	"testing"
	"time"

	"github.com/TinyMurky/go-clock/internal/config"
	"github.com/TinyMurky/go-clock/internal/service/clockhands"
	"github.com/TinyMurky/go-clock/internal/testhelpers"
)

func TestSecondHand(t *testing.T) {
	testCases := []struct {
		now   time.Time
		point clockhands.Point
	}{
		{
			testhelpers.SetTime(t, 0, 0, 0), // 應該要是朝正上方
			clockhands.Point{
				X: config.ClockRadius,
				Y: config.ClockRadius - config.SecondHandLength,
			},
		},
	}

	for _, testCase := range testCases {
		got := clockhands.SecondHand(testCase.now)

		testhelpers.RoughlyEqualFloat32(t, got.X, testCase.point.X)
		testhelpers.RoughlyEqualFloat32(t, got.Y, testCase.point.Y)

	}
}
