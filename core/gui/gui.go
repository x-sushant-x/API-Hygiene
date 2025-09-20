package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/x-sushant-x/API-Hygiene/core"
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

func buildRightPanel() (*fyne.Container, *widget.Label) {
	statusCodeCheck := widget.NewLabel("Status Code Check")

	statusCodeCheckResponse := widget.NewLabel("Pending...")

	checks := container.NewVBox(
		widget.NewLabel("Test Results:"),
		container.NewHBox(
			statusCodeCheck,
			statusCodeCheckResponse,
		),
	)

	return container.NewStack(checks), statusCodeCheckResponse
}

func buildMainContent() *container.Split {
	methodSelect := buildMethodSelector()
	endpointEntry := buildEndpointEntry()
	bodyEntry := buildBodyEntry()

	rightPanel, statusCodeLabel := buildRightPanel()

	submitBtn := widget.NewButton("Run Tests", func() {
		method := methodSelect.Selected
		endpoint := endpointEntry.Text

		runner := core.NewHygieneRunner(endpoint, method)
		report := runner.CheckHygiene()

		var resultText string
		if report.ErrorMessage != "" {
			resultText = fmt.Sprintf("Error: %s", report.ErrorMessage)

			rightPanel.Objects = []fyne.CanvasObject{
				widget.NewLabel(resultText),
			}
			rightPanel.Refresh()

		} else {
			if report.StatusCode.IsValidCode {
				statusCodeLabel.SetText("Valid")
				statusCodeLabel.TextStyle = fyne.TextStyle{Bold: true}

				statusCodeLabel.Refresh()
				statusCodeLabel.Alignment = fyne.TextAlignLeading
			} else {
				statusCodeLabel.SetText("Invalid")
				statusCodeLabel.TextStyle = fyne.TextStyle{Bold: true}
				statusCodeLabel.Refresh()
				statusCodeLabel.Alignment = fyne.TextAlignLeading
			}

			statusCodeLabel.Refresh()
		}
	})

	leftPanel := container.NewPadded(container.NewVBox(
		widget.NewLabel("Method:"),
		methodSelect,
		widget.NewLabel("Endpoint:"),
		endpointEntry,
		widget.NewLabel("Request Body:"),
		bodyEntry,
		layout.NewSpacer(),
		submitBtn,
	))
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
