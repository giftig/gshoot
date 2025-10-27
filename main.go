package main

import (
	"fmt"
	"image"
	"image/color"
	"log/slog"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/kbinani/screenshot"
)

var translucentColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0xbb}
var borderColor = color.White
var textColor = color.White

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	a := app.New()
	win := a.NewWindow("Test")
	img := getScreenshot(0)
	selector := getSelector()

	content := container.NewWithoutLayout(img, selector)

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
	win.SetContent(content)
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

func getSelector() *fyne.Container {
	rect := canvas.NewRectangle(translucentColor)
	rect.StrokeColor = borderColor
	rect.StrokeWidth = 1
	rect.SetMinSize(fyne.NewSize(200, 200))

	txt := canvas.NewText("0 x 0", textColor)

	selector := container.New(layout.NewCenterLayout(), rect, txt)
	// TODO: Temporary to show the rendering. These should start at 0 and hidden, and will be resized
	// in response to mouse drag events on a custom widget
	selector.Move(fyne.NewPos(100, 100))
	selector.Resize(fyne.NewSize(200, 200))

	return selector
}
