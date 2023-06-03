package layout

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func GridWrapLayout() {
	a := app.New()
	w := a.NewWindow("Grid Wrap Layout")
	t1 := canvas.NewText("1", color.White)
	t2 := canvas.NewText("2", color.White)
	t3 := canvas.NewText("3", color.White)
	c := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 50)),
		t1, t2, t3)
	w.SetContent(c)
	w.ShowAndRun()
}
