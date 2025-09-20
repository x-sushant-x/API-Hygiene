package core

import (
	"net/http"
)

var returnCodesList = map[string][]int{
	"GET":    {200},
	"POST":   {201},
	"PUT":    {200},
	"DELETE": {200},
	"PATCH":  {200},
}

func checkStatusCode(requestType string, resp *http.Response) StatusCode {
	validCodes, _ := returnCodesList[requestType]

	statusCode := StatusCode{}

	statusCode.Got = resp.StatusCode

	for _, validCode := range validCodes {
		if statusCode.Got == validCode {
			statusCode.IsValidCode = true
		}
	}

	return statusCode
}
