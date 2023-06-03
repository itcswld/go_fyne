package layout

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func CenterLayout() {
	a := app.New()
	w := a.NewWindow("Center")
	w.Resize(fyne.NewSize(200, 200))
	img := canvas.NewImageFromResource(theme.FyneLogo())
	img.FillMode = canvas.ImageFillOriginal
	// img.FillMode = canvas.ImageFillContain //doesn't work
	txt := canvas.NewText("Overlay", color.White)
	c := container.New(layout.NewCenterLayout(), img, txt)

	w.SetContent(c)
	w.ShowAndRun()
}
