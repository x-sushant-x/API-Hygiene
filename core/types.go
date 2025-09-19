package core

type HygieneReport struct {
	StatusCode StatusCode
}

type StatusCode struct {
	ValidCodes  []int
	Got         int
	IsValidCode bool
	IsErrorCode bool
}
