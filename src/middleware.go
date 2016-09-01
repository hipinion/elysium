package elysium

import (
	"fmt"
	"net/http"
)

func sessionMiddleware(next http.Handler) http.Handler {
	fmt.Println("um?")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		cookie, _ := r.Cookie("elysium_sid")
		fmt.Println("Cookie: ", cookie)
		fmt.Println("guh?")
		next.ServeHTTP(w, r)
	})
}
