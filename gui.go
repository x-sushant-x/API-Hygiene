package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func buildMethodSelector() *widget.Select {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	methodSelect := widget.NewSelect(methods, func(value string) {})
	methodSelect.SetSelected("GET")
	return methodSelect
}

func buildEndpointEntry() *widget.Entry {
	endpointEntry := widget.NewEntry()
	endpointEntry.SetPlaceHolder("Enter API endpoint...")
	return endpointEntry
}

func buildBodyEntry() *widget.Entry {
	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder("Enter request body (for POST/PUT)...")
	return bodyEntry
}

func buildSubmitButton() *widget.Button {
	return widget.NewButton("Run Tests", func() {
		// TODO: Add request execution logic here
	})
}

func buildLeftPanel() *fyne.Container {
	methodSelect := buildMethodSelector()
	endpointEntry := buildEndpointEntry()
	bodyEntry := buildBodyEntry()
	submitBtn := buildSubmitButton()

	return container.NewPadded(container.NewVBox(
		widget.NewLabel("Method:"),
		methodSelect,
		widget.NewLabel("Endpoint:"),
		endpointEntry,
		widget.NewLabel("Request Body:"),
		bodyEntry,
		layout.NewSpacer(),
		submitBtn,
	))
}

func buildMainContent() *container.Split {
	leftPanel := buildLeftPanel()
	rightPanel := container.NewStack() // Placeholder for response/output panel
	mainContent := container.NewHSplit(leftPanel, rightPanel)
	mainContent.Offset = 0.35
	return mainContent
}

func StartApplication() {
	myApp := app.New()
	myWindow := myApp.NewWindow("API Hygiene")

	mainContent := buildMainContent()
	myWindow.SetContent(mainContent)
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}
