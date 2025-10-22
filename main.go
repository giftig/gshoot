package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Test")

	fmt.Println("Hmm")
	label := widget.NewLabel("Hello world")
	box := container.NewVBox(label)

	win.SetFullScreen(true)
	win.SetContent(box)
	win.ShowAndRun()
}
