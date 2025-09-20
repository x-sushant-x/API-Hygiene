package core

type HygieneReport struct {
	StatusCode StatusCode
}

type StatusCode struct {
	Got         int
	IsValidCode bool
}
