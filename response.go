package http_json_rpc

import (
	"strconv"
	"log"
)

type JsonRpcResponse struct {
	Jsonrpc   string
	Error     map[string]interface{}
	Id        int64
	Result    interface{}
	plainText string
}

func (r *JsonRpcResponse)GetInt() (int64) {
	var result int64
	switch unknown := r.Result.(type) {
	default:
		result = 0
	case bool:
		if unknown {
			result = 1
		} else {
			result = 0
		}
	case string:

		if intval, err := strconv.ParseInt(unknown, 10, 64); err == nil {
			result = intval
		} else if floatval, err := strconv.ParseFloat(unknown, 64); err == nil {
			result = int64(floatval)
		}
	case int:
		result = int64(unknown)
	case float32:
	case float64:
		result = int64(unknown)
	}
	return result
}
func (r *JsonRpcResponse)GetString() (string) {
	var result string
	switch unknown := r.Result.(type) {
	default:
		result = ""
	case bool:
		result = strconv.FormatBool(unknown)
	case int:
		result = strconv.FormatInt(int64(unknown), 2)
	case float32:
		result = strconv.FormatFloat(float64(unknown), 'f', -1, 32)
	case float64:
		result = strconv.FormatFloat(unknown, 'f', -1, 64)
	case string:
		result = unknown
	}
	return result
}
func (r *JsonRpcResponse)GetMap() (map[string]interface{}) {
	if result, ok := r.Result.(map[string]interface{}); ok {
		return result
	} else {
		log.Fatal(r.Result)
		return nil
	}
}
func (r *JsonRpcResponse)GetArray() ([]interface{}) {
	if result, ok := r.Result.([]interface{}); ok {
		return result
	} else {
		return nil
	}
}
