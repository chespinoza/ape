package ape

import (
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
func New() (s *Server) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s.Router = httprouter.New()
	return
}

//RunHTTP put the http server to run, the function don't return
//if return is only in case of error
func (s *Server) RunHTTP(addr string, readTimeout, writeTimeout int) error {
	handler := context.ClearHandler(s.Router)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Duration(readTimeout),
		WriteTimeout: time.Duration(writeTimeout),
		Handler:      handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
