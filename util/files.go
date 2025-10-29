package util

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

// Generate a filename for a screenshot. This will include the current datetime and a short uuid
func ScreenshotFilename(clock Clock) string {
	u := uuid.New().String()[:7]
	t := clock.Now().Format("2006-01-02_15-04-05")

	return fmt.Sprintf("screenshot_%s_%s.png", t, u)
}

// Return the full path to a destination for the screenshot. Will default to ~/Desktop or use
// the GSHOOT_SCREENSHOT_DIR env var as destination directory
func ScreenshotDir() (string, error) {
	p := os.Getenv("GSHOOT_SCREENSHOT_DIR")
	if p != "" {
		return p, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, "Desktop"), nil
}

func ScreenshotPath(clock Clock) (string, error) {
	dir, err := ScreenshotDir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, ScreenshotFilename(clock)), nil
}
