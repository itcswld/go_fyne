package layout

import (
	"fmt"
	widgets "test/ch1_widgets/cus_widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Tab() {
	a := app.New()
	w := a.NewWindow("Tab")
	w.Resize(fyne.NewSize(300, 200))

	tabs := container.NewAppTabs(
		container.NewTabItem("Home", widget.NewLabel("Hello, \n  welcome home!")),
		container.NewTabItem("Product", widget.NewLabel("Go fyne the world!")),
	)

	txtField_acc, e_acc := widgets.TextField("Acc")
	txtField_pwd, e_pwd := widgets.TextField("Pwd")
	btn_submit := widget.NewButton("Submit", func() {
		fmt.Println("acc :", e_acc.Text, "\n", "pwd :", e_pwd.Text)
	})
	loginPage := container.New(layout.NewVBoxLayout(), txtField_acc, txtField_pwd, btn_submit)

	tabs.Append(container.NewTabItemWithIcon("Login", theme.LoginIcon(), loginPage))

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.ShowAndRun()
}
