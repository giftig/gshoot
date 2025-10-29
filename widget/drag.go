package widget

import "fyne.io/fyne/v2"

// fyne's drag events are focused on dragging a component around, and its mouse events don't
// include tracking whether it's a "drag", so this stored details of whether we're mid-drag
type DragEvent struct {
	Start   *fyne.PointEvent
	Current *fyne.PointEvent
}
