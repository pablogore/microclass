package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	hello "github.com/bketelsen/microclass/module1/gomicro/greeter/api/rpc/proto/hello"
)

func main() {
	req, err := proto.Marshal(&hello.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := http.Post("http://localhost:8080/greeter/hello", "application/protobuf", bytes.NewReader(req))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	rsp := &hello.Response{}
	if err := proto.Unmarshal(b, rsp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}
