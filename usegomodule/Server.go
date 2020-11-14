package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/wangyuche/practice/creategomodule/testmath"
)

func httpHandle(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello fasthttp")
}

func main() {
	testmath.Add(1, 2)
	if err := fasthttp.ListenAndServe("0.0.0.0:8080", httpHandle); err != nil {
		fmt.Println("start fasthttp fail:", err.Error())
	}
}
