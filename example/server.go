package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	kong "github.com/giovanni-liboni/kongo-jwt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/matryer/respond"
	"github.com/spf13/viper"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func HandleAuthEndpoint(w http.ResponseWriter, r *http.Request) {
	// Retrive auth user (if any)
	user := context.Get(r, "auth")
	respond.With(w, r, http.StatusOK, user)
}

// HandleLogin releases the token
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Authenticate your users before call GetToken method
	token, err := kong.GetToken("test", "123")
	if err != nil {
		respond.With(w, r, http.StatusInternalServerError, err.Error())
	} else {
		respond.With(w, r, http.StatusOK, TokenAuthentication{Token: token})
	}
	return
}

func main() {
	// User viper to configure the app
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/endpoint_auth", HandleAuthEndpoint).Methods("GET")
	router.HandleFunc("/login", HandleLogin).Methods("POST")

	n := negroni.New(kong.AuthMiddleware())
	n.UseHandler(router)

	log.Println("Web server started on port 8080")

	http.ListenAndServe(":8080", n)
}
