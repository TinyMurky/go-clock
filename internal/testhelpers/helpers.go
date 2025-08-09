// Plackage testhelpers has some helpper function
package testhelpers

import (
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
