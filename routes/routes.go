package routes

import (
	"log"
	"net/http"

	"github.com/Robson019/api-rest-go/controllers"
	"github.com/Robson019/api-rest-go/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeMiddleware)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalities", controllers.FetchPersonalities).Methods("Get")
	router.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	router.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	router.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
