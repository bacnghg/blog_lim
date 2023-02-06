package common

type resultResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

type successResponsePaging struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    resultResponse `json:"data"`
}

type successResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data, paging, filter interface{}) *successResponsePaging {
	result := resultResponse{Data: data, Paging: paging, Filter: filter}
	return &successResponsePaging{Code: 200, Message: "success", Data: result}

}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{Code: 200, Message: "success", Data: data}
}

func SimpleErrorResponse(statusCode int, message string) *successResponse {
	return &successResponse{Code: statusCode, Message: message}
}
