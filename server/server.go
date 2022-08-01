// app.go

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jorgesiachoque08/melicoupons/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

//function in charge of initializing the server
func (s *Server) Initialize() {
	s.Router = mux.NewRouter()
	s.initializeRoutes()

}

//function in charge of running the server
func (s *Server) Run(serverPort string) {
	fmt.Println("Run server: http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, s.Router))
}

//function in charge of initializing the router
func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/coupon", handlers.Coupon).Methods("POST")
	s.Router.HandleFunc("/topFavorites", handlers.TopFavorites).Methods("GET")
}
