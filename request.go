package http_json_rpc

import (
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
)

type JsonRpcPayload struct {
	Method  string                           `json:"method"`
	Params  map[string]interface{}           `json:"params"`
	Jsonrpc string                           `json:"jsonrpc"`
	Id      int                        	 `json:"id"`
}

func NewJsonRpcPayload(method string, params map[string]interface{}) (*JsonRpcPayload) {

	payload := JsonRpcPayload{
		Method:method,
		Params:params,
		Jsonrpc:"2.0",
	}
	return &payload
}

func Call(host, methodName string, params map[string]interface{}) (*JsonRpcResponse,error) {
	var response *JsonRpcResponse
	var err error
	payload := NewJsonRpcPayload(
		methodName,
		params,
	)
	jsonStr, err := json.Marshal(payload)
	if err != nil {
		err=err
	} else {
		resp, err := http.Post(host, "application/x-www-form-urlencoded", strings.NewReader(string(jsonStr)))
		defer resp.Body.Close()
		if err != nil {
			err=err
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err!=nil{
				err=err
			}else{
				response = &JsonRpcResponse{}
				if err:=json.Unmarshal(body, response);err==nil{
					err=err
				}
				response.plainText=string(body)
			}
		}
	}
	return response,err
}