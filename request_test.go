package http_json_rpc

import (
	"testing"
	"log"
	"fmt"
)

func TestCall(t *testing.T) {
	host:="http://d2d.med.docker/rpc/order"
	params:=map[string]interface{}{
		"orderId":70,
	}
	response,err:= Call(host,"get",params)
	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println(response.GetMap()["doctorTitleName"])
		fmt.Println(response.plainText)
	}
}
