package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	srv *http.Server
}

func New(port string) *mux.Router {
	s := &server{}
	router := mux.NewRouter()
	router.HandleFunc("/health", s.health).Methods("GET")
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}
	s.srv = srv
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err.Error())
		}
	}()
	go func() {
		s.signal()
	}()
	return router
}

func (this *server) signal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("%s\n", "Server Shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := this.srv.Shutdown(ctx); err != nil {
		panic(err.Error())
	}
}

func (this *server) health(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}
