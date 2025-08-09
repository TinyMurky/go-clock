// Package testhelpers has some helpper function
package testhelpers

import (
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/TinyMurky/go-clock/internal/config"
)

// SetTime return time.Time that is base on 1970/1/1 with hour minute second
func SetTime(t testing.TB, hour, minute, second int) time.Time {
	t.Helper()

	timezone, _ := time.LoadLocation(config.Timezone)
	return time.Date(1970, time.January, 1, hour, minute, second, 0, timezone)
}

// RoughlyEqualFloat32 will pass if the different between
// got and want is less than 1e-7
func RoughlyEqualFloat32(t testing.TB, got, want float32) {
	t.Helper()

	const threshold = 1e-7

	if math.Abs(float64(got-want)) > threshold {
		t.Errorf("expect \"%v\", got \"%v\"", want, got)
	}
}

// AssertEqual use reflect equal to determine equalTestHand
func AssertEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect \"%v\", got \"%v\"", want, got)
	}
}
