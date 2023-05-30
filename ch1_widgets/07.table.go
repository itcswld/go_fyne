package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Table() {
	var data = [][]string{
		{"col1", "col2"},
		{"row2_col1", "row2_col2"},
	}

	a := app.New()
	w := a.NewWindow("Table")
	w.Resize(fyne.NewSize(300, 300))

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("row2_col1")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[tci.Row][tci.Col])
		})

	w.SetContent(table)
	w.ShowAndRun()
}
