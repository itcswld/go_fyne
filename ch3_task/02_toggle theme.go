package task

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func DarkTheme() {

	tm := "light"
	os.Setenv("FYNE_THEME", tm)
	a := app.New()

	w := a.NewWindow("Toggle Theme")
	w.Resize(fyne.NewSize(200, 200))
	btn := widget.NewButton("toggle theme", func() {
		if tm == "light" {
			a.Settings().SetTheme(theme.DarkTheme())
			tm = "dark"
		} else {
			a.Settings().SetTheme(theme.LightTheme())
			tm = "light"
		}

	})
	btn.Resize(fyne.NewSize(180, 30))

	w.SetContent(container.NewWithoutLayout(btn))
	w.ShowAndRun()
}
