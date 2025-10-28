package main

import (
	"fmt"
	"image"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/kbinani/screenshot"

	"github.com/giftig/gshoot/widget"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	runApp()
}

func runApp() {
	a := app.New()
	win := a.NewWindow("Test")
	img := getScreenshot(0)
	w := widget.NewScreenshotWidget(img)

	win.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "C" {
			slog.Info(fmt.Sprintf("FYNE_SCALE = %s", os.Getenv("FYNE_SCALE")))
			slog.Info(fmt.Sprintf("Window dimensions: %s", win.Canvas().Size()))
			slog.Info(fmt.Sprintf("Scale: %f", win.Canvas().Scale()))
			return
		}

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

func getScreenshot(displayNum int) *canvas.Image {
	scr, err := screenshot.CaptureRect(screenshot.GetDisplayBounds(displayNum))
	if err != nil {
		slog.Error("Failed to take screenshot")
		os.Exit(1)
	}

	img := image.Image(scr)
	bounds := img.Bounds().Max
	cmp := canvas.NewImageFromImage(img)
	cmp.Resize(fyne.NewSize(float32(bounds.X), float32(bounds.Y)))

	return cmp
}
