package modules

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)
	// convert the id type from string to int
	email := params["email"]
	// create an empty user of type models.User
	var u_ins users
	// decode the json request to user
	json.NewDecoder(r.Body).Decode(&u_ins)
	print(u_ins.FullName)
	/*if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}*/
	// call update user to update the user
	updatedRows := updateUser(email, u_ins)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
	// format the response message
	var response = JsonResponse{
		Type:    "success",
		Message: msg}

	// send the response
	json.NewEncoder(w).Encode(response)
}
func updateUser(email string, u_ins users) int64 {
	// create the postgres db connection
	db := setupDB()
	// close the db connection
	defer db.Close()
	// create the update sql query
	sqlStatement := "UPDATE Signupinv SET full_name=$2, num=$3, promo_code=$4 WHERE email=$1;"
	res, err := db.Exec(sqlStatement, email, u_ins.FullName, u_ins.PhoneNumber, u_ins.PromoCode)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
