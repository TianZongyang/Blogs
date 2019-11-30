package response

const (
	NO_AUTH       = 400
	NO_AUTH_MSG   = "No Auth"
	EXCEPTION     = 500
	EXCEPTION_MSG = "System error"
)

type ResponseInfo struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *ResponseInfo {
	return &ResponseInfo{
		Success: true,
		Code:    0,
		Msg:     "",
		Data:    data,
	}
}

func SuccessWithCode(data interface{}, code int, msg string) *ResponseInfo {
	return &ResponseInfo{
		Success: true,
		Code:    code,
		Msg:     msg,
		Data:    data,
	}
}

func Error(data interface{}) *ResponseInfo {
	return &ResponseInfo{
		Success: false,
		Code:    EXCEPTION,
		Msg:     EXCEPTION_MSG,
		Data:    nil,
	}
}

func NoAuth() *ResponseInfo {
	return &ResponseInfo{
		Success: false,
		Code:    NO_AUTH,
		Msg:     NO_AUTH_MSG,
		Data:    nil,
	}
}
