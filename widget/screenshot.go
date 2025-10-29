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

	Image    image.Image
	Selector *SelectorWidget

	OnCapture func(image.Image)
}

func NewScreenshotWidget(img image.Image, onCapture func(image.Image)) *ScreenshotWidget {
	w := &ScreenshotWidget{
		Image:     img,
		Selector:  NewSelectorWidget(),
		OnCapture: onCapture,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *ScreenshotWidget) CreateRenderer() fyne.WidgetRenderer {
	bounds := w.Image.Bounds().Max
	img := canvas.NewImageFromImage(w.Image)
	img.Resize(fyne.NewSize(float32(bounds.X), float32(bounds.Y)))

	content := container.NewWithoutLayout(img, w.Selector)
	return widget.NewSimpleRenderer(content)
}

func (w *ScreenshotWidget) MouseDown(e *desktop.MouseEvent) {
	w.Selector.SetOrigin(fyne.NewPos(e.AbsolutePosition.X, e.AbsolutePosition.Y))
}
func (w *ScreenshotWidget) MouseUp(e *desktop.MouseEvent) {
	if w.Selector.Origin != nil && w.Selector.Dimensions != nil {
		w.capture(*w.Selector.Origin, *w.Selector.Dimensions)
	}

	w.Selector.Origin = nil
	w.Selector.Dimensions = nil
	w.Selector.Hide()
}
func (w *ScreenshotWidget) MouseMoved(e *desktop.MouseEvent) {
	if w.Selector.Origin == nil {
		return
	}
	width := e.AbsolutePosition.X - w.Selector.Origin.X
	height := e.AbsolutePosition.Y - w.Selector.Origin.Y

	w.Selector.SetDimensions(fyne.NewSize(width, height))
}
func (w *ScreenshotWidget) MouseIn(e *desktop.MouseEvent) {}
func (w *ScreenshotWidget) MouseOut()                     {}

func (w *ScreenshotWidget) Cursor() desktop.Cursor {
	return desktop.CrosshairCursor
}

// Commit the screenshot selection, cropping the full screenshot image down to the specified area
// Call the OnCapture function with the result
func (w *ScreenshotWidget) capture(origin fyne.Position, dim fyne.Size) {
	// SubImage isn't part of the Image interface so it needs to be type-asserted to an interface
	// containing SubImage. In practice this is almost all types of image.
	subimage := w.Image.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	cropped := subimage.SubImage(
		image.Rect(int(origin.X), int(origin.Y), int(origin.X+dim.Width), int(origin.Y+dim.Height)),
	)
	w.OnCapture(cropped)
}

func NewSelectorWidget() *SelectorWidget {
	rect := canvas.NewRectangle(translucentColor)
	rect.StrokeColor = borderColor
	rect.StrokeWidth = 1

	label := canvas.NewText("0 x 0", textColor)

	w := &SelectorWidget{
		Rect:       rect,
		Label:      label,
		Origin:     nil,
		Dimensions: nil,
	}
	w.ExtendBaseWidget(w)

	return w
}

func (w *SelectorWidget) CreateRenderer() fyne.WidgetRenderer {
	content := container.New(layout.NewCenterLayout(), w.Rect, w.Label)

	return widget.NewSimpleRenderer(content)
}

func (w *SelectorWidget) SetOrigin(origin fyne.Position) {
	w.Origin = &origin
	w.Move(*w.Origin)
}

func (w *SelectorWidget) SetDimensions(dimensions fyne.Size) {
	w.Dimensions = &dimensions
	w.Resize(*w.Dimensions)
	w.Rect.SetMinSize(*w.Dimensions)
	w.Label.Text = fmt.Sprintf("%d x %d", int32(w.Dimensions.Width), int32(w.Dimensions.Height))
	w.RefreshVisibility()
}

// Set to visible if the selector has origin and dimensions, otherwise set invisible
func (w *SelectorWidget) RefreshVisibility() {
	if w.Origin != nil && w.Dimensions != nil && w.Dimensions.Width > 0 && w.Dimensions.Height > 0 {
		w.Show()
	} else {
		w.Hide()
	}
}
