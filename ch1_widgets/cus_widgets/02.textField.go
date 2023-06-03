package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func TextField(title string) (*fyne.Container, *widget.Entry) {

	label := widget.NewLabel(title)
	input := widget.NewEntry()
	input.Resize(fyne.NewSize(250, 30))
	label_input := container.New(layout.NewFormLayout(), label, input)

	return label_input, input
}
