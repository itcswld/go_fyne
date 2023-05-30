package widgets

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//To create a number input spinner similar to Bootstrap 5
func InputSpinner() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Number Spinner")
	//display the number
	numberEntry := widget.NewEntry()
	numberEntry.SetText("0")
	numberEntry.Disable()
	//increment button
	incrementButton := widget.NewButton("+", func() {
		currentValue, _ := strconv.Atoi(numberEntry.Text)
		currentValue++
		numberEntry.SetText(strconv.Itoa(currentValue))
	})
	//decrement button
	decrementButton := widget.NewButton("-", func() {
		currentValue, _ := strconv.Atoi(numberEntry.Text)
		currentValue--
		numberEntry.SetText(strconv.Itoa(currentValue))
	})

	content := container.NewHBox(
		decrementButton,
		numberEntry,
		incrementButton,
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
