package layout

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func GridLayout() {
	a := app.New()
	w := a.NewWindow("Grid")
	w.Resize(fyne.NewSize(200, 200))

	txt1 := canvas.NewText("txt1", color.White)
	txt2 := canvas.NewText("txt2", color.White)
	txt3 := canvas.NewText("txt3", color.White)

	grid := container.New(layout.NewGridLayout(2), txt1, txt2, txt3)
	w.SetContent(grid)
	w.ShowAndRun()
}
