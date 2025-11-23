package widget

import (
	"fmt"
	"image"
	"image/color"
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/giftig/gshoot/config"
	"github.com/giftig/gshoot/math"
)

var translucentColor = color.NRGBA{R: 0, G: 0, B: 0, A: 0xbb}
var borderColor = color.White
var textColor = color.White

type SelectorWidget struct {
	widget.BaseWidget

	Rect  *canvas.Rectangle
	Label *canvas.Text

	Origin *fyne.Position
	Dest   *fyne.Position
}

type ScreenshotWidget struct {
	widget.BaseWidget

	Image    image.Image
	Selector *SelectorWidget

	OnCapture  func(image.Image, config.EditConfig)
	EditConfig config.EditConfig
}

func NewScreenshotWidget(
	img image.Image,
	onCapture func(image.Image, config.EditConfig),
) *ScreenshotWidget {
	w := &ScreenshotWidget{
		Image:      img,
		Selector:   NewSelectorWidget(),
		OnCapture:  onCapture,
		EditConfig: config.EditConfig{PostEdit: false},
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
	if e.Button != desktop.MouseButtonPrimary {
		return
	}
	w.Selector.SetOrigin(fyne.NewPos(e.AbsolutePosition.X, e.AbsolutePosition.Y))
}
func (w *ScreenshotWidget) MouseUp(e *desktop.MouseEvent) {
	if e.Button != desktop.MouseButtonPrimary {
		return
	}
	if w.Selector.Origin != nil && w.Selector.Dest != nil {
		w.capture(*w.Selector.Origin, *w.Selector.Dest)
	}

	w.Selector.Origin = nil
	w.Selector.Dest = nil
	w.Selector.RefreshDisplay()
}
func (w *ScreenshotWidget) MouseMoved(e *desktop.MouseEvent) {
	if w.Selector.Origin == nil {
		return
	}

	w.Selector.SetDest(fyne.NewPos(e.AbsolutePosition.X, e.AbsolutePosition.Y))
}
func (w *ScreenshotWidget) MouseIn(e *desktop.MouseEvent) {}
func (w *ScreenshotWidget) MouseOut()                     {}

// Alt-drag triggers opening the screenshot in an image editor after capture
func (w *ScreenshotWidget) KeyDown(e *fyne.KeyEvent) {
	slog.Error("WTF")
	fmt.Printf("Key down: %s", e.Name)
	if e.Name == "alt" {
		w.EditConfig.PostEdit = true
	}
}
func (w *ScreenshotWidget) KeyUp(e *fyne.KeyEvent) {
	if e.Name == "alt" {
		w.EditConfig.PostEdit = false
	}
}

func (w *ScreenshotWidget) FocusGained()     {}
func (w *ScreenshotWidget) FocusLost()       {}
func (w *ScreenshotWidget) TypedRune(r rune) {}
func (w *ScreenshotWidget) TypedKey(e *fyne.KeyEvent) {
	fmt.Printf("Key typed: %s", e.Name)
}

func (w *ScreenshotWidget) Cursor() desktop.Cursor {
	return desktop.CrosshairCursor
}

// Commit the screenshot selection, cropping the full screenshot image down to the specified area
// Call the OnCapture function with the result
func (w *ScreenshotWidget) capture(origin fyne.Position, dest fyne.Position) {
	// SubImage isn't part of the Image interface so it needs to be type-asserted to an interface
	// containing SubImage. In practice this is almost all types of image.
	subimage := w.Image.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	cropped := subimage.SubImage(image.Rect(int(origin.X), int(origin.Y), int(dest.X), int(dest.Y)))
	w.OnCapture(cropped, w.EditConfig)
}

func NewSelectorWidget() *SelectorWidget {
	rect := canvas.NewRectangle(translucentColor)
	rect.StrokeColor = borderColor
	rect.StrokeWidth = 1

	label := canvas.NewText("0 x 0", textColor)

	w := &SelectorWidget{
		Rect:   rect,
		Label:  label,
		Origin: nil,
		Dest:   nil,
	}
	w.ExtendBaseWidget(w)

	return w
}

func (w *SelectorWidget) CreateRenderer() fyne.WidgetRenderer {
	content := container.New(layout.NewCenterLayout(), w.Rect, w.Label)

	return widget.NewSimpleRenderer(content)
}

func (w *SelectorWidget) GetBounds() (fyne.Position, fyne.Size) {
	if w.Origin == nil || w.Dest == nil {
		return fyne.NewPos(0, 0), fyne.NewSize(0, 0)
	}

	x := math.Min(w.Origin.X, w.Dest.X)
	y := math.Min(w.Origin.Y, w.Dest.Y)
	width := math.Abs(w.Origin.X - w.Dest.X)
	height := math.Abs(w.Origin.Y - w.Dest.Y)

	return fyne.NewPos(x, y), fyne.NewSize(width, height)
}

func (w *SelectorWidget) SetOrigin(origin fyne.Position) {
	w.Origin = &origin
	w.RefreshDisplay()
}

func (w *SelectorWidget) SetDest(dest fyne.Position) {
	w.Dest = &dest
	w.RefreshDisplay()
}

// Recalculate dimensions based on updated origin or destination
func (w *SelectorWidget) RefreshDisplay() {
	pos, size := w.GetBounds()
	w.Move(pos)
	w.Resize(size)
	w.Rect.SetMinSize(size)
	w.Label.Text = fmt.Sprintf("%d x %d", int32(size.Width), int32(size.Height))

	if size.Width != 0 && size.Height != 0 {
		w.Show()
	} else {
		w.Hide()
	}
}
