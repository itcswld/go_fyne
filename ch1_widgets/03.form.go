package widgets

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Form() {
	a := app.New()
	w := a.NewWindow("Form")
	w.Resize(fyne.NewSize(300, 300))

	e_text := widget.NewEntry()
	mle_text := widget.NewMultiLineEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{{
			Text:   "Title",
			Widget: e_text,
		}},
		OnSubmit: func() {
			log.Println("text:", e_text.Text)
			log.Println("textArea:", mle_text.Text)
			w.Close()
		},
	}
	//append items
	form.Append("Content", mle_text)

	w.SetContent(form)
	w.ShowAndRun()
}
