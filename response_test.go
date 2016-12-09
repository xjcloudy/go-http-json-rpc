package http_json_rpc

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestJsonRpcResponse_GetString(t *testing.T) {
	r := JsonRpcResponse{
		Result:1.20001,
	}
	fmt.Println(r.GetString())
}
func TestJsonRpcResponse_GetInt(t *testing.T) {
	r := JsonRpcResponse{
		Result:false,
	}
	fmt.Println(r.GetInt())
}
func TestJsonRpcResponse_GetArray(t *testing.T) {
	r:=JsonRpcResponse{
		Result:map[string]interface{}{
			"a":map[string]string{
				"b":"c",
			},
		},
	}
	fmt.Println(r.GetMap()["a"])
}
func TestJsonRpcResponse_GetMap(t *testing.T) {

	r:=JsonRpcResponse{}
	jsondata:=`{"result":[{"a":1},1]}`
	json.Unmarshal([]byte(jsondata),&r)
	fmt.Println(r.GetArray())
}
