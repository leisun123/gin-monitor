package model

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error"` // 错误项
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}
