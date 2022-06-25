package routes

import (
	"log"
	"net/http"

	"github.com/felipejzsouza/go-api-rest/controllers"
	"github.com/felipejzsouza/go-api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.BuscarPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidade/{id}", controllers.BuscarPersonalidade).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidade", controllers.CadastrarPersonalidade).Methods(http.MethodPost)
	r.HandleFunc("/api/personalidade/{id}", controllers.DeletarPersonalidade).Methods(http.MethodDelete)
	r.HandleFunc("/api/personalidade/{id}", controllers.EditarPersonalidade).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
