package layout

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func MaxLayout() {
	a := app.New()
	w := a.NewWindow("Max")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	txt := canvas.NewText("Overlay", color.White)
	//sets all items in the container to be the same size as the container
	c := container.New(layout.NewMaxLayout(), img, txt)

	w.SetContent(c)
	w.ShowAndRun()
}
