package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "0712"
	DB_NAME     = "Signup"
)

// DB set up
func setupDB() *sql.DB {
	var db *sql.DB

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

type users struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int    `json:"num"`
	PromoCode   string `json:"promo_code"`
	Reference   string `json:"refer"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []users `json:"data"`
	Message string  `json:"message"`
}

// Main function
func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/getusers/", GetUser).Methods("GET")

	// Create a movie
	router.HandleFunc("/createusers/", CreateUsers).Methods("POST")

	//Delete a user
	router.HandleFunc("/users/{email}", DeleteUser).Methods("DELETE")

	//Update a user
	router.HandleFunc("/getupdate/{email}", UpdateUser).Methods("PUT")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Get all users

// response and request handlers
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting users...")

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

// Create a movie

// response and request handlers
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	full_name := r.FormValue("fullname")
	email := r.FormValue("eemail")
	password := r.FormValue("pwd")
	num := r.FormValue("number")
	promo_code := r.FormValue("code")
	reff := r.FormValue("referal")

	var response = JsonResponse{}

	if full_name == "" || email == "" || password == "" || num == "" || promo_code == "" || reff == "" {
		response = JsonResponse{Type: "error", Message: "You are missing any of the parameter."}
	} else {
		db := setupDB()

		printMessage("Creating users..")

		fmt.Println("Inserting new user with details: " + full_name + " with mail " + email + "with password" + password + "with phone number" + num + "promocode given by" + promo_code + "referred by" + reff)

		var lastInsertID string
		err := db.QueryRow("INSERT INTO Signupinv(full_name, email, password, num, promo_code, refer) VALUES($1, $2, $3, $4, $5, $6) returning full_name;", full_name, email, password, num, promo_code, reff).Scan(&lastInsertID)
		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The details has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	email := params["email"]

	var response = JsonResponse{}

	if email == "" {
		response = JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting user from DB")

		_, err := db.Exec("DELETE FROM Signupinv where email = $1", email)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

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

func createConnection() {
	panic("unimplemented")
}

/*func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	email := params["email"]

	var response = JsonResponse{}

	if email == "" {
		response = JsonResponse{Type: "error", Message: "You are missing users parameter."}
	} else {
		db := setupDB()

		_, err := db.Exec("UPDATE Signupinv SET fullname=$2, num=$3, promo_code=$4 WHERE email=$1; returning *", email)

		printMessage("Updating user from DB")

		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The user has been updateed successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}
*/
