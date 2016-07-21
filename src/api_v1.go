package elysium

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var ()

func API_Page() {

}

func API_Topic() {

}

func API_Forum() {

}

func API_User() {

}

func API_v1_UsersHandler(w http.ResponseWriter, r *http.Request) {
	var us Users
	var wheres []string
	var whereVals []interface{}

	var getUsername = r.URL.Query().Get("user_name")
	if getUsername != "" {
		wheres = append(wheres, "u.user_name = ?")
		whereVals = append(whereVals, getUsername)
	}

	var getEmail = r.URL.Query().Get("user_email")
	if getEmail != "" {
		wheres = append(wheres, "u.user_email = ?")
		whereVals = append(whereVals, getEmail)
	}

	whereString := strings.Join(wheres, " OR ")

	if len(wheres) > 0 {
		whereString = " WHERE " + whereString
	}

	users, err := DB.Query("SELECT u.user_name, u.user_id, u.user_email FROM users u "+whereString+" LIMIT 25", whereVals...)

	for users.Next() {

		var u User
		err = users.Scan(&u.Name, &u.ID, &u.Email)

		if err != nil {
		}

		us.Users = append(us.Users, u)

	}

	userJSON, _ := json.Marshal(us)
	fmt.Fprintln(w, string(userJSON))
}

func API_v1_TopicsHandler(w http.ResponseWriter, r *http.Request) {

}

func API_v1_ForumsHandler(w http.ResponseWriter, r *http.Request) {

}

func API_v1_PostsHandler(w http.ResponseWriter, r *http.Request) {

}
