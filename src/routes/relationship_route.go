package routes

import (
	"ca-server/src/controllers"

	"github.com/gorilla/mux"
)

func SetupRelationshipRoutes(api *mux.Router) {
	api.HandleFunc("/friend-request", controllers.CreateFriendRequest).Methods("POST")
	api.HandleFunc("/accept-request", controllers.AcceptFriendRequest).Methods("POST")

}
