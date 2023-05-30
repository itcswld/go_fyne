package task

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

//selecting a file and then retrieve the file path
func GetFilePath() {
	a := app.New()
	w := a.NewWindow("Entry")
	w.Resize(fyne.NewSize(300, 300))

	lb := widget.NewLabel("Selected File Path: ")

	btn := widget.NewButton("Select File", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err == nil && uri != nil {
				lb.SetText("Selected File Path: " + uri.URI().Path())
				// Send the selected file path to the backend
				fmt.Println(uri.URI().Path())
			}
		}, w)
	})

	w.SetContent(container.NewVBox(lb, btn))

	w.ShowAndRun()
}
