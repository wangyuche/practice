module gomodule

go 1.15

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/wangyuche/practice/creategomodule v0.0.0-20201114084807-90574950b9d9
	github.com/wangyuche/practice/usegomodule/src/httpserver v0.0.0
)

replace github.com/wangyuche/practice/usegomodule/src/httpserver => ./src/httpserver
