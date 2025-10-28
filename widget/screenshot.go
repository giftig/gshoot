package widget

import (
	"fmt"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var translucentColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0xbb}
var borderColor = color.White
var textColor = color.White

type SelectorWidget struct {
	widget.BaseWidget

	Rect  *canvas.Rectangle
	Label *canvas.Text

	Origin     *fyne.Position
	Dimensions *fyne.Size
}

type ScreenshotWidget struct {
	widget.BaseWidget

	Image    *canvas.Image
	Selector *SelectorWidget
	Origin   *image.Point
}

func NewScreenshotWidget(img *canvas.Image) *ScreenshotWidget {
	w := &ScreenshotWidget{
		Image:    img,
		Selector: NewSelectorWidget(),
		Origin:   nil,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *ScreenshotWidget) CreateRenderer() fyne.WidgetRenderer {
	content := container.NewWithoutLayout(w.Image, w.Selector)
	return widget.NewSimpleRenderer(content)
}

func (w *ScreenshotWidget) MouseDown(e *desktop.MouseEvent) {
	fmt.Println("MOUSE DOWN")
}
func (w *ScreenshotWidget) MouseUp(e *desktop.MouseEvent) {
	fmt.Println("MOUSE UP")
}

func NewSelectorWidget() *SelectorWidget {
	rect := canvas.NewRectangle(translucentColor)
	rect.StrokeColor = borderColor
	rect.StrokeWidth = 1

	label := canvas.NewText("0 x 0", textColor)
	origin := fyne.NewPos(100, 100)
	dim := fyne.NewSize(300, 300)

	w := &SelectorWidget{
		Rect:  rect,
		Label: label,
		// TODO: Temporary to show the rendering. These should start as nil instead
		Origin:     &origin,
		Dimensions: &dim,
	}
	w.ExtendBaseWidget(w)

	return w
}

func (w *SelectorWidget) CreateRenderer() fyne.WidgetRenderer {
	content := container.New(layout.NewCenterLayout(), w.Rect, w.Label)

	if w.Origin != nil && w.Size != nil {
		w.Move(*w.Origin)
		w.Resize(*w.Dimensions)
		w.Rect.SetMinSize(*w.Dimensions)
	} else {
		w.Hide()
	}
	return widget.NewSimpleRenderer(content)
}
