package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Label_Entry_Btn() {
	a := app.New()
	w := a.NewWindow("Barcdoe Printer")
	w.Resize(fyne.NewSize(400, 400))
	//label
	l_code := widget.NewLabel("Barcdoe")
	//--position
	l_code.Move(fyne.NewPos(40, 20))
	//input
	e_code := widget.NewEntry()
	e_code.SetPlaceHolder("Enter you barcode")
	e_code.Resize(fyne.NewSize(250, 30))
	e_code.Move(fyne.NewPos(40, 50))
	//button
	b_submit := widget.NewButton("submit", func() {
		fmt.Println(e_code.Text)
		e_code.SetText("")
	})
	b_submit.Resize(fyne.NewSize(250, 30))
	b_submit.Move(fyne.NewPos(40, 90))
	//w.SetContent(container.NewVBox(l_code, e_code, b_submit))
	w.SetContent(container.NewWithoutLayout(l_code, e_code, b_submit))

	w.ShowAndRun()
}
