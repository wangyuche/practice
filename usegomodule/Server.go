package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/wangyuche/practice/creategomodule"
	"github.com/wangyuche/practice/usegomodule/src/httpserver"
)

func main() {
	x := creategomodule.Add(1, 1)
	fmt.Printf("x: %d\n", x)
	router := httpserver.New(os.Getenv("port"))
	router.HandleFunc("/test", test).Methods("GET")
	quit := make(chan os.Signal, 1)
	<-quit
}

func test(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello word"))
}
