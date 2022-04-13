package apimodels

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
