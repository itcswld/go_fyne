package widgets

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Silder() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Number Slider")

	numberEntry := widget.NewEntry()
	numberEntry.SetText("0")

	numberSlider := widget.NewSlider(0, 100)
	numberSlider.Step = 1
	numberSlider.OnChanged = func(value float64) {
		numberEntry.SetText(strconv.Itoa(int(value)))
	}

	numberEntry.OnChanged = func(value string) {
		if num, err := strconv.Atoi(value); err == nil {
			numberSlider.Value = float64(num)
		}
	}

	selectedNumberLabel := widget.NewLabel("Selected number: 0")

	numberEntry.OnChanged = func(value string) {
		selectedNumberLabel.SetText(fmt.Sprintf("Selected number: %s", value))
	}

	content := container.NewVBox(
		numberEntry,
		numberSlider,
		selectedNumberLabel,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
