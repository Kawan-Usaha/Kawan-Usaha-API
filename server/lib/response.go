package lib

func Ok(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}
}

func Error(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
		"data":    data,
	}
}
