package ape

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

const DefaultADDR = ":8080"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Basic Test Done!!\n")
}
func TestBasicApe(t *testing.T) {
	server := New()
	server.Router.GET("/", Index)
	fmt.Printf("Running server at %v\n", DefaultADDR)
	server.RunHTTP(DefaultADDR, 3600, 3600)
}
