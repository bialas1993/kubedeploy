package http

import (
	"errors"
)

const (
	UnchangedStatus  = 409
	ApiFailureStatus = "Failure"
)

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func ValidResponse(err error, response *Response) error {
	if response.Code == UnchangedStatus {
		return nil
	}

	if response.Status == ApiFailureStatus {
		return errors.New("http: Request failed.")
	}

	return err
}
