package api

import (
	"encoding/json"
)

type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type ErrorResponse struct {
	ErrorCode int                    `json:"error_code"`
	Message   string                 `json:"message"`
	ErrorInfo map[string]interface{} `json:"error_info"`
}
