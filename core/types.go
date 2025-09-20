package core

import "time"

type HygieneReport struct {
	ResponseBody string
	ResponseTime time.Duration
	ErrorMessage string
	StatusCode   StatusCode
}

type StatusCode struct {
	Got         int
	IsValidCode bool
}
