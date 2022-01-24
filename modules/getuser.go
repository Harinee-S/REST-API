package modules

import (
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	//printMessage("Getting users...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM Signupinv")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var details []users

	// Foreach movie
	for rows.Next() {
		var full_name string
		var email string
		var password string
		var num int
		var promo_code string
		var refer string

		err = rows.Scan(&full_name, &email, &password, &num, &promo_code, &refer)

		// check errors
		checkErr(err)

		details = append(details, users{FullName: full_name, Email: email, Password: password, PhoneNumber: num, PromoCode: promo_code, Reference: refer})
	}

	var response = JsonResponse{Type: "success", Data: details}

	json.NewEncoder(w).Encode(response)
}
