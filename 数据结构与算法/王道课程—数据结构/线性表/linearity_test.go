package linearity

import "testing"

var (
	linearity = NewLinearity(100)
)

func TestLinearity_AddElement(t *testing.T) {
	t.Run("linearity_append", func(t *testing.T) {
		linearity.AddElement(1)
	})
}
