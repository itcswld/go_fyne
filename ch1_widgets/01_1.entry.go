package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Entry() {
	a := app.New()
	w := a.NewWindow("Entry")
	w.Resize(fyne.NewSize(300, 300))
	e := widget.NewEntry()
	btn := widget.NewButton("submit", func() {
		fmt.Println(e.Text)
		e.SetText("")
		//how to set focus to the widget enty filed
		w.Canvas().Focus(e)
	})
	w.SetContent(container.NewVBox(e, btn))

	w.ShowAndRun()
}
