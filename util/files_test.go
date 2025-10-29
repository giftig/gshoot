package util

import (
	"os"
	"path"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScreenshotFilename(t *testing.T) {
	clock := MockClock(time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC))

	expected := regexp.MustCompile("screenshot_2020-01-02_03-04-05_[0-9a-f]+.png")
	actual := ScreenshotFilename(clock)

	assert.Regexp(t, expected, actual, "filename should match pattern")
}

func TestScreenshotDirEnvVar(t *testing.T) {
	t.Setenv("GSHOOT_SCREENSHOT_DIR", "/foo/bar")

	res, err := ScreenshotDir()
	assert.Nil(t, err)
	assert.Equal(t, res, "/foo/bar")
}

func TestScreenshotDirDefault(t *testing.T) {
	assert := assert.New(t)
	t.Setenv("GSHOOT_SCREENSHOT_DIR", "")
	t.Setenv("HOME", "/home/testuser")

	home, err := os.UserHomeDir()
	assert.Nil(err)

	expected := path.Join(home, "Desktop")

	actual, err := ScreenshotDir()
	assert.Nil(err)
	assert.Equal(actual, expected)
}
