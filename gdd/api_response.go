package gdd

import "net/http"

type ApiResponse struct {
	Status int
	Msg    string
	Data   interface{}
}

func OK(data interface{}) *ApiResponse {
	return &ApiResponse{
		Status: http.StatusOK,
		Msg:    http.StatusText(http.StatusOK),
		Data:   data,
	}
}

func ERR(data interface{}) *ApiResponse {
	return &ApiResponse{
		Status: http.StatusInternalServerError,
		Msg:    http.StatusText(http.StatusInternalServerError),
		Data:   data,
	}
}

func (a *ApiResponse) SetStatus(status int) *ApiResponse {
	a.Status = status
	a.Msg = http.StatusText(status)
	return a
}

func (a *ApiResponse) SetMsg(msg string) *ApiResponse {
	a.Msg = msg
	return a
}

func (a *ApiResponse) SetData(data interface{}) *ApiResponse {
	a.Data = data
	return a
}
