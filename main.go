package main

import (
	"image"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kbinani/screenshot"

	"github.com/giftig/gshoot/util"
	"github.com/giftig/gshoot/widget"
	"github.com/giftig/gshoot/writer"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	runApp()
}

func runApp() {
	a := app.New()
	os.Setenv("FYNE_DISABLE_DPI_DETECTION", "1")
	clock := util.NewClock()

	win := a.NewWindow("gshoot")
	img := getScreenshot(0)
	w := widget.NewScreenshotWidget(img, func(captured image.Image) {
		slog.Info("Capture obtained")
		win.Close()
		err := writer.WriteScreenshot(captured, clock)

		if err != nil {
			slog.Error("Failed to write screenshot to file", slog.Any("err", err))
		}
	})

	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "Escape" {
			win.Close()
			return
		}
	})

	win.SetPadded(false)
	win.SetFullScreen(true)
	win.SetContent(w)
	win.ShowAndRun()
}

func getScreenshot(displayNum int) image.Image {
	scr, err := screenshot.CaptureRect(screenshot.GetDisplayBounds(displayNum))
	if err != nil {
		slog.Error("Failed to take screenshot")
		os.Exit(1)
	}

	return image.Image(scr)
}
