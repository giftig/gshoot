package writer

import (
	"image"
	"image/png"
	"os"

	"github.com/giftig/gshoot/util"
)

func WriteScreenshot(img image.Image, clock util.Clock) (string, error) {
	path, err := util.ScreenshotPath(clock)
	if err != nil {
		return "", err
	}

	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return "", err
	}

	return path, nil
}
