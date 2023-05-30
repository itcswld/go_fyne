package widgets

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ProgressBar() {
	a := app.New()
	w := a.NewWindow("ProgressBar")
	w.Resize(fyne.NewSize(300, 300))
	pb := widget.NewProgressBar()
	pbi := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i < 10.0; i += 1.0 {
			time.Sleep(time.Millisecond * 250)
			pb.SetValue(i)
		}
	}()

	w.SetContent(container.NewVBox(pb, pbi))
	w.ShowAndRun()
}
