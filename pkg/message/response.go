package message

func SuccessResponseSlice(data *[]any) map[string]interface{} {
	return map[string]interface{}{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func SuccessResponse(data any) map[string]interface{} {
	return map[string]interface{}{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func ErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
