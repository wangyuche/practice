package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/wangyuche/practice/creategomodule"
)

func httpHandle(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello fasthttp")
}

func main() {
	x := creategomodule.Add(1, 1)
	fmt.Printf("x: %d\n", x)
	if err := fasthttp.ListenAndServe("0.0.0.0:8080", httpHandle); err != nil {
		fmt.Println("start fasthttp fail:", err.Error())
	}
}
