package spotify

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func ReqAuth(w http.ResponseWriter, r *http.Request) {
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Fprintln(w, "User already authenticated: ", u)
	}

	gothic.BeginAuthHandler(w, r)
}

func GetAccessTok(w http.ResponseWriter, r *http.Request) LogIn {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println("Error getting access token")
		fmt.Fprintln(w, err)
	}

	fmt.Printf("User log-in: %s\n", user.Name)
	fmt.Printf("Token expiry date: %s\n", user.ExpiresAt)

	newLogIn := LogIn{
		Name:        user.Name,
		AccessToken: user.AccessToken,
		ExpiryDate:  user.ExpiresAt.String(),
	}
	return newLogIn
}
