package widgets

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Choices() {
	a := app.New()
	w := a.NewWindow("Choices")
	ck := widget.NewCheck("Option1", func(b bool) {
		log.Println("Check set to", b)
	})
	rd := widget.NewRadioGroup([]string{"option1", "option2"}, func(s string) {
		log.Println("Radio set to", s)
	})
	s := widget.NewSelect([]string{"option1", "option2"}, func(s string) {
		log.Panicln("Select set to", s)
	})

	w.SetContent(container.NewVBox(ck, rd, s))

	w.ShowAndRun()
}
