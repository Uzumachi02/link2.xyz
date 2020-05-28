package utils

// APIResponse
type APIResponse struct {
	Error    ErrorResponse `json:"error,omitempty"`
	Response interface{}   `json:"response,omitempty"`
}

// APIResponseItems
type APIResponseItems struct {
	Error    ErrorResponse `json:"error,omitempty"`
	Response response      `json:"response,omitempty"`
}

type response struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}

// ErrorResponse ...
type ErrorResponse struct {
	ErrorCode     int         `json:"error_code,omitempty"`
	ErrorMsg      string      `json:"error_msg,omitempty"`
	RequestParams interface{} `json:"request_params,omitempty"`
}
