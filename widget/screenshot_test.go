package widget

import (
	"testing"

	"fyne.io/fyne/v2/driver/desktop"
)

func TestScreenshotWidgetMouseable(t *testing.T) {
	var _ desktop.Mouseable = (*ScreenshotWidget)(nil)
}

func TestScreenshotWidgetKeyable(t *testing.T) {
	var _ desktop.Keyable = (*ScreenshotWidget)(nil)
}
