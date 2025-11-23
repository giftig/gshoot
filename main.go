package main

import (
	"image"
	"log/slog"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/jasonlovesdoggo/displayindex"
	"github.com/kbinani/screenshot"

	"github.com/giftig/gshoot/config"
	"github.com/giftig/gshoot/edit"
	"github.com/giftig/gshoot/util"
	"github.com/giftig/gshoot/widget"
	"github.com/giftig/gshoot/writer"
)

var clock util.Clock = util.NewClock()

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	runApp()
}

func runApp() {
	a := app.New()
	clock := util.NewClock()
	os.Setenv("FYNE_DISABLE_DPI_DETECTION", "1")

	win := a.NewWindow("gshoot")
	img := getScreenshot()
	w := widget.NewScreenshotWidget(
		img,
		func(captured image.Image, cfg config.EditConfig) {
			slog.Info("Capture obtained", slog.Any("cfg", cfg))
			win.Close()

			f, err := writer.WriteScreenshot(captured, clock)
			if err != nil {
				slog.Error("Failed to write screenshot to file", slog.Any("err", err))
				return
			}

			if cfg.PostEdit {
				edit.EditScreenshot(f)
			}
		},
		func() {
			win.Close()
		},
	)

	win.SetPadded(false)
	win.SetFullScreen(true)
	win.SetContent(w)
	win.Canvas().Focus(w)
	win.ShowAndRun()
}

func getScreenshot() image.Image {
	displayNum, err := displayindex.CurrentDisplayIndex()
	if err != nil {
		slog.Error("Could not find display index; defaulting to display 0")
		displayNum = 0
	}

	scr, err := screenshot.CaptureRect(screenshot.GetDisplayBounds(displayNum))
	if err != nil {
		slog.Error("Failed to take screenshot")
		os.Exit(1)
	}

	return image.Image(scr)
}
