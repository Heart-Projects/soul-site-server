package model

type ResponseData struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *ResponseData) S() *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: "操作成功",
		Code:    200,
	}
	return res
}

func (r *ResponseData) SuccessWithMsg(msg string) *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: msg,
		Code:    200,
	}
	return res
}

func (r *ResponseData) SuccessWithCode(code int) *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: "操作成功",
		Code:    code,
	}
	return res
}

func (r *ResponseData) SuccessWithData(data interface{}) ResponseData {
	res := ResponseData{
		Success: true,
		Message: "操作成功",
		Code:    200,
		Data:    data,
	}
	return res
}

func (r *ResponseData) SuccessWithCodeData(code int, data interface{}) *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: "操作成功",
		Code:    code,
		Data:    data,
	}
	return res
}

func (r *ResponseData) SuccessWithMsgData(msg string, data interface{}) *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: msg,
		Code:    200,
		Data:    data,
	}
	return res
}

func (r *ResponseData) SuccessWithCodeMsg(code int, msg string) *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: msg,
		Code:    code,
	}
	return res
}

func (r *ResponseData) SUCCESS(code int, data interface{}, msg string) *ResponseData {
	res := &ResponseData{
		Success: true,
		Message: msg,
		Code:    code,
		Data:    data,
	}
	return res
}

func (r *ResponseData) F() *ResponseData {
	res := &ResponseData{
		Message: "操作失败",
		Code:    500,
	}
	return res
}

func (r *ResponseData) FailWithMsg(msg string) *ResponseData {
	res := &ResponseData{
		Message: msg,
		Code:    500,
	}
	return res
}

func (r *ResponseData) FailWithCode(code int) *ResponseData {
	res := &ResponseData{
		Message: "操作失败",
		Code:    code,
	}
	return res
}

func (r *ResponseData) FailWithData(data interface{}) *ResponseData {
	res := &ResponseData{
		Message: "操作失败",
		Code:    500,
		Data:    data,
	}
	return res
}

func (r *ResponseData) FailWithCodeData(code int, data interface{}) *ResponseData {
	res := &ResponseData{
		Message: "操作失败",
		Code:    code,
		Data:    data,
	}
	return res
}

func (r *ResponseData) FailWithMsgData(msg string, data interface{}) *ResponseData {
	res := &ResponseData{
		Message: msg,
		Code:    500,
		Data:    data,
	}
	return res
}

func (r *ResponseData) FailWithCodeMsg(code int, msg string) *ResponseData {
	res := &ResponseData{
		Message: msg,
		Code:    code,
	}
	return res
}

func (r *ResponseData) FAIL(code int, data interface{}, msg string) *ResponseData {
	res := &ResponseData{
		Message: msg,
		Code:    code,
		Data:    data,
	}
	return res
}

// 导出全局变量
var Response = &ResponseData{}
