package core

import (
	"errors"
	"net/http"
)

var returnCodesList = map[string][]int{
	"GET":    {200, 304, 404, 500},
	"POST":   {201, 400, 409, 500},
	"PUT":    {200, 204, 400, 404, 500},
	"DELETE": {200, 204, 404, 500},
	"PATCH":  {200, 204, 400, 404, 500},
}

func CheckStatusCode(requestType string, resp http.Response) (StatusCode, error) {
	validCodes, found := returnCodesList[requestType]

	if !found {
		return StatusCode{}, errors.New("unknown or un-supported request type")
	}

	statusCode := StatusCode{}

	statusCode.Got = resp.StatusCode
	statusCode.ValidCodes = validCodes

	for _, validCode := range validCodes {
		if statusCode.Got == validCode {
			statusCode.IsValidCode = true
		}
	}

	return statusCode, nil
}
