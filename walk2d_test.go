package walk2d

import (
	"testing"
)

func TestHLine(t *testing.T) {
	HLine(0, 10, 5, func(x, y int) bool {
		return false
	})
}
