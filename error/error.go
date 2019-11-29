package error

import "fmt"

func DefineError(prefix string, index int, msg string) func(error) map[string]interface{} {
	return func(err error) map[string]interface{} {
		return map[string]interface{}{
			"code":    fmt.Sprintf("%s-%03d", prefix, index),
			"message": msg,
			"cause":   err.Error(),
		}
	}
}
