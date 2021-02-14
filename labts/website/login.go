package website

import "net/http"

var authToken string

//SetAuth sets the auth token
func SetAuth(auth string) {
	authToken = auth
}

func login() {
	loginPage := http.Get(LabTSURL)

}
