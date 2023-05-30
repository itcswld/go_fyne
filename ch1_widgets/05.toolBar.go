package widgets

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Toolbar() {
	a := app.New()
	w := a.NewWindow("Toolbar")

	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Panicln("New Doc..")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Panicln("Display help")
		}),
	)
	content := container.NewBorder(tb, nil, nil, nil, widget.NewLabel("Content"))

	w.SetContent(content)
	w.ShowAndRun()
}
