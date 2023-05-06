package response

import (
	"reflect"

	"github.com/mbaraa/apollo-music/errors"
)

type responseStatus string

const (
	ok    responseStatus = "ok"
	error responseStatus = "error"
)

type responseTemplate struct {
	Status     responseStatus   `json:"status"`
	HttpStatus int              `json:"-"`
	Body       any              `json:"body"`
	ErrorCode  errors.ErrorCode `json:"errorCode"`
	ErrorMsg   string           `json:"errorMsg"`
}

type json map[string]any

// BuildResponse accepts an error code and a body,
// and returns a json response and an http status code, following this template
//
// {status: "ok" | "error", body?: {}, errorCode: number, errorMsg: string}
//
// where a `body` will only exist if the response is without any errors,
// and `errorCode` & `errorMsg` will exist only on errors
func BuildResponse(errCode errors.ErrorCode, body any) (json, int) {
	resp := map[string]any{
		"status": ok,
	}

	if errCode != errors.None {
		resp["status"] = error
		resp["errorCode"] = errCode
		resp["errorMsg"] = errCode.String()
		return resp, errCode.StatusCode()
	}

	if reflect.TypeOf(body).Kind() == reflect.Struct {
		resp["body"] = body
	}
	return resp, 200
}
