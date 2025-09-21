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
	selectBox := widget.NewSelect(methods, nil)
	selectBox.SetSelected("GET")
	return selectBox
}

func buildEndpointEntry() *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter API endpoint...")
	return entry
}

func buildBodyEntry() *widget.Entry {
	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Enter request body (for POST/PUT)...")
	return entry
}

type RightPanel struct {
	Container *fyne.Container
	labels    map[string]*widget.Label
}

func newRightPanel() *RightPanel {
	rp := &RightPanel{
		labels: make(map[string]*widget.Label),
	}

	rp.addCheck(API_VERSIONING)
	rp.addCheck(ENDPOINT_NAMING_CONVENTION)
	rp.addCheck(AUTHORIZATION_CHECK)
	rp.addCheck(API_VERSIONING)
	rp.addCheck(IS_HTTPS_CONFIGURED)
	rp.addCheck(PAGINATION)
	rp.addCheck(REQUEST_CONSISTENCY)
	rp.addCheck(RESPONSE_CONSISTENCY)

	rp.Container = container.NewVBox(
		widget.NewLabel("Test Results:"),
	)

	for name, label := range rp.labels {
		rp.Container.Add(container.NewHBox(
			widget.NewLabel(name+":"),
			layout.NewSpacer(),
			label,
		))
	}

	return rp
}

func (rp *RightPanel) addCheck(name string) {
	rp.labels[name] = widget.NewLabel("Pending...")
}

func (rp *RightPanel) updateCheck(name, value string, isValid bool) {
	if label, ok := rp.labels[name]; ok {
		label.SetText(value)
		label.TextStyle = fyne.TextStyle{Bold: true}
		if isValid {
			label.Alignment = fyne.TextAlignLeading
		} else {
			label.Alignment = fyne.TextAlignLeading
		}
		label.Refresh()
	}
}

func buildMainContent() *container.Split {
	methodSelect := buildMethodSelector()
	endpointEntry := buildEndpointEntry()
	bodyEntry := buildBodyEntry()
	rightPanel := newRightPanel()

	submitBtn := widget.NewButton("Run Tests", func() {
		method := methodSelect.Selected
		endpoint := endpointEntry.Text

		runner := core.NewHygieneRunner(endpoint, method)
		report := runner.CheckHygiene()

		if report.ErrorMessage != "" {
			rightPanel.updateCheck("Status Code", fmt.Sprintf("Error: %s", report.ErrorMessage), false)
			return
		}

		if report.StatusCode.IsValidCode {
			rightPanel.updateCheck("Status Code", "Valid", true)
		} else {
			rightPanel.updateCheck("Status Code", "Invalid", false)
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

	mainContent := container.NewHSplit(leftPanel, rightPanel.Container)
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
