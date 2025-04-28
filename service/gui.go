package service

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Window() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Fyne")

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Welcome!"),
		widget.NewButton("Click me", func() {
			println("Button clicked!")
		}),
	))

	myWindow.ShowAndRun()
}
