package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/marcosmourabro/go-rest-api/controllers"
	"github.com/marcosmourabro/go-rest-api/middlewere"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewere.ContetTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonalities).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonalities).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonalities).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"*"}))(r)))
}
