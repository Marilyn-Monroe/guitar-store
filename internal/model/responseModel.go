package model

type ResponseModel struct {
	Errors *[]struct {
		Code    *int32  `json:"code,omitempty"`
		Field   *string `json:"field,omitempty"`
		Message *string `json:"message,omitempty"`
	} `json:"errors,omitempty"`
	Message *string `json:"message,omitempty"`
}
