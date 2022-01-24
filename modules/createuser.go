package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	full_name := r.FormValue("fullname")
	email := r.FormValue("eemail")
	password := r.FormValue("pword")
	num := r.FormValue("number")
	promo_code := r.FormValue("code")
	reff := r.FormValue("referal")

	fmt.Println(full_name, email, password, num, promo_code, reff)

	var response = JsonResponse{}

	if full_name == "" || email == "" || password == "" || num == "" || promo_code == "" || reff == "" {
		response = JsonResponse{Type: "error", Message: "You are missing any of the parameter."}
	} else {
		db := setupDB()

		//printMessage("Creating users..")

		fmt.Println("Inserting new user with details: " + full_name + " with mail " + email + "with password" + password + "with phone number" + num + "promocode given by" + promo_code + "referred by" + reff)

		var lastInsertID string
		err := db.QueryRow("INSERT INTO Signupinv(full_name, email, password, num, promo_code, refer) VALUES($1, $2, $3, $4, $5, $6) returning full_name;", full_name, email, password, num, promo_code, reff).Scan(&lastInsertID)
		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The details has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
