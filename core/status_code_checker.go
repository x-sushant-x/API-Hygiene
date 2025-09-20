package core

import (
	"errors"
	"net/http"
)

var returnCodesList = map[string][]int{
	"GET":    {200},
	"POST":   {201},
	"PUT":    {200},
	"DELETE": {200},
	"PATCH":  {200},
}

func CheckStatusCode(requestType string, resp http.Response) (StatusCode, error) {
	validCodes, found := returnCodesList[requestType]

	if !found {
		return StatusCode{}, errors.New("unknown or un-supported request type")
	}

	statusCode := StatusCode{}

	statusCode.Got = resp.StatusCode

	for _, validCode := range validCodes {
		if statusCode.Got == validCode {
			statusCode.IsValidCode = true
		}
	}

	return statusCode, nil
}
