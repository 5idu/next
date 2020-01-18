package ecode

// Response response
type Response struct {
	Code    APICode     `json:"code,omitempty"`    // 响应码
	Message string      `json:"message,omitempty"` // 前台展示信息
	Detail  string      `json:"detail,omitempty"`  // 堆栈信息
	Data    interface{} `json:"data,omitempty"`    // 响应数据
}

// NewResponse new response
func NewResponse(code APICode, err error, data ...interface{}) *Response {
	resp := Response{}
	if err != nil {
		resp.Detail = err.Error()
	}
	resp.Code = code
	resp.Message = APICodeMapZH[code]
	if len(data) > 0 {
		resp.Data = data[0]
	}
	return &resp
}
