package modules

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	email := params["email"]

	var response = JsonResponse{}

	if email == "" {
		response = JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
	} else {
		db := setupDB()

		//printMessage("Deleting user from DB")

		_, err := db.Exec("DELETE FROM Signupinv where email = $1", email)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
