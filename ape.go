package ape

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

//Server type is an instance of httprouter
type Server struct {
	Router *httprouter.Router
}

//New Create a new Ape Server
func New() *Server {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := &Server{}
	s.Router = httprouter.New()
	return s
}

//RunHTTP startthe http server
//addr is the TCP address to listen on.
//readTimeout is the max duration before timing out read of the request in secs
//writeTimeout is the max duration before timing out write of the response in secs
func (s *Server) RunHTTP(addr string, readTimeout, writeTimeout int) {
	handler := context.ClearHandler(s.Router)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		Handler:      handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
