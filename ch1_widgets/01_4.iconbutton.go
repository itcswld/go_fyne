package widgets

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func init() {
	os.Setenv("FYNE_FONT", "TaipeiSansTCBeta-Bold.ttf")
}

func IconBtn() {

	os.Setenv("FYNE_FONT", "TaipeiSansTCBeta-Bold.ttf")

	a := app.New()
	w := a.NewWindow("Icon Button")
	w.Resize(fyne.NewSize(200, 200))
	btn := widget.NewButtonWithIcon("Home 家", theme.HomeIcon(), func() {
		fmt.Println("welcome home.")
	})
	btn.Resize(fyne.NewSize(150, 30))
	w.SetContent(container.NewWithoutLayout(btn))
	w.ShowAndRun()
}
