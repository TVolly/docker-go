package responses

func SerializeData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data": data,
	}
}
