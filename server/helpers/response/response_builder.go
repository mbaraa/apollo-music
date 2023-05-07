package response

import (
	"reflect"

	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
)

type responseStatus string

const (
	ok    responseStatus = "ok"
	error responseStatus = "error"
)

type responseTemplate struct {
	Status    responseStatus   `json:"status"`
	Data      any              `json:"data"`
	ErrorCode errors.ErrorCode `json:"errorCode"`
	ErrorMsg  string           `json:"errorMsg"`
}

// Build accepts an error code and a data,
// and returns a json response and an http status code, following this template
//
// {status: "ok" | "error", data?: {}, errorCode: number, errorMsg: string}
//
// where a `data` will only exist if the response is without any errors,
// and `errorCode` & `errorMsg` will exist only on errors
func Build(errCode errors.ErrorCode, data any) (entities.JSON, int) {
	resp := map[string]any{
		"status": ok,
	}

	if errCode != errors.None {
		resp["status"] = error
		resp["errorCode"] = errCode
		resp["errorMsg"] = errCode.String()
		return resp, errCode.StatusCode()
	}

	dataKind := reflect.TypeOf(data).Kind()
	if dataKind == reflect.Struct || dataKind == reflect.Map {
		resp["data"] = data
	}
	return resp, 200
}
