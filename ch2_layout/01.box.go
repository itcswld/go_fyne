package layout

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BoxLayout() {
	a := app.New()
	w := a.NewWindow("Box")
	w.Resize(fyne.NewSize(300, 300))
	t_code := canvas.NewText("Code", color.White)
	e_code := widget.NewEntry()
	e_code.SetPlaceHolder("Enter you barcode")
	e_code.Resize(fyne.NewSize(200, 30)) //-- size doesnt work
	contentH := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(), t_code, e_code, layout.NewSpacer(), layout.NewSpacer())

	btn := widget.NewButton("submit", func() {
		fmt.Println(e_code.Text)
	})
	//layout.NewSpacer() between will center all widgets
	center := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btn, layout.NewSpacer())
	w.SetContent(container.New(layout.NewVBoxLayout(), contentH, center))
	w.ShowAndRun()
}
