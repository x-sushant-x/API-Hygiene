package gui

import (
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/x-sushant-x/API-Hygiene/core"
)

var selectedRequestMethod = "GET"

func buildMethodSelector(bodyEntry *widget.Entry) *widget.Select {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	selectBox := widget.NewSelect(methods, func(s string) {
		if s == http.MethodPost || s == http.MethodPut || s == http.MethodPatch {
			bodyEntry.Enable()
		} else {
			bodyEntry.Disable()
		}
	})
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

	rp.addCheck(STATUS_CODE_CHECK)
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

func buildMainContent(requestMethod string, rootWindow fyne.Window) *container.Split {
	bodyEntry := buildBodyEntry()

	methodSelect := buildMethodSelector(bodyEntry)
	endpointEntry := buildEndpointEntry()

	if requestMethod != http.MethodPost && requestMethod != http.MethodPut && requestMethod != http.MethodPatch {
		bodyEntry.Disable()
	}

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

	uploadSpecificationFile := widget.NewButtonWithIcon("Upload OpenAPI or Swagger", theme.UploadIcon(), func() {

		fileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				if reader == nil {
					fmt.Println("No file selected")
					return
				}
				fmt.Println("Selected file:", reader.URI().Path())
				reader.Close()
			}, rootWindow)

		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml", ".json"}))

		fileDialog.Show()
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
		uploadSpecificationFile,
	))

	mainContent := container.NewHSplit(leftPanel, rightPanel.Container)
	mainContent.Offset = 0.35
	return mainContent
}

func StartApplication() {
	myApp := app.NewWithID("com.sushant.api_hygiene")
	myWindow := myApp.NewWindow("API Hygiene")

	mainContent := buildMainContent(selectedRequestMethod, myWindow)
	myWindow.SetContent(mainContent)
	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}
