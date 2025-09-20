package core

import (
	"net/http"
	"time"
)

type HygieneRunner struct {
	endpoint      string
	requestMethod string
}

func NewHygieneRunner(endpoint string, requestMethod string) HygieneRunner {
	return HygieneRunner{
		endpoint:      endpoint,
		requestMethod: requestMethod,
	}
}

func (hg HygieneRunner) CheckHygiene() HygieneReport {
	report := HygieneReport{}

	resp, err := hg.hitAPI()

	if err != nil {
		return HygieneReport{
			ErrorMessage: err.Error(),
		}
	}

	statusCode := checkStatusCode(hg.requestMethod, resp)

	report.StatusCode = statusCode

	return report
}

func (hg HygieneRunner) hitAPI() (*http.Response, error) {
	req, err := http.NewRequest(hg.requestMethod, hg.endpoint, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
