package main

//importing neccessary libararies
import (
	"Project2/modules"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/getusers/", modules.GetUser).Methods("GET")

	// Create a movie
	router.HandleFunc("/createusers/", modules.CreateUsers).Methods("POST")

	//Delete a user
	router.HandleFunc("/deleteuser/{email}", modules.DeleteUser).Methods("DELETE")

	//Update a user
	router.HandleFunc("/updateuser/{email}", modules.UpdateUser).Methods("PUT")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))

}
