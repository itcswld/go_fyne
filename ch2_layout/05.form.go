package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func FormLayout() {
	a := app.New()
	w := a.NewWindow("Form")
	w.Resize(fyne.NewSize(250, 30))
	l_code := widget.NewLabel("code")
	e_code := widget.NewEntry()
	e_code.Resize(fyne.NewSize(250, 30))
	label_input := container.New(layout.NewFormLayout(), l_code, e_code)

	w.SetContent(label_input)
	w.ShowAndRun()

}
