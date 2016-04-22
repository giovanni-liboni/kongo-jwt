package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	kong "github.com/giovanni-liboni/kongo-jwt"
	"github.com/gorilla/mux"
	"github.com/matryer/respond"
	"github.com/spf13/viper"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

// Autenticazione

// Chiamata ad url accessibile solo da utenti autenticati

// Chiamata ad url accessibile a tutti

func HandleAuthEndpoint(w http.ResponseWriter, r *http.Request) {

}
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Authenticate your users before call GetToken method
	token, err := kong.GetToken("test", "1")
	if err != nil {
		respond.With(w, r, http.StatusInternalServerError, err.Error())
		return
	} else {
		respond.With(w, r, http.StatusOK, TokenAuthentication{Token: token})
	}
}

func main() {
	// Use this path to find configuration path
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/endpoint_auth", HandleAuthEndpoint).Methods("GET")
	router.HandleFunc("/auth", HandleAuthEndpoint).Methods("POST")
	router.HandleFunc("/login", HandleLogin).Methods("POST")

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(router)

	log.Println("Web server started on port 8080")

	http.ListenAndServe(":8080", n)
}
