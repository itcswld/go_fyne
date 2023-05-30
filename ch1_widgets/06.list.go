package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func List() {

	var data = []string{"Apple", "Banana", "Cat"}

	a := app.New()
	w := a.NewWindow("List")
	w.Resize(fyne.NewSize(300, 300))

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[lii])
		})

	w.SetContent(list)
	w.ShowAndRun()
}
