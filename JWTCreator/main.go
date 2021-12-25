package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySignKey = []byte(os.Getenv("SECRET_KEY"))

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	log.Println("Valid Token: ", validToken)
	if err != nil {
		fmt.Println("Failed to generate token: ", err)
	}
	fmt.Fprintf(w, string(validToken))
}

func handleRequests() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "JonathanBrull"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(mySignKey)
	if err != nil {
		fmt.Println("Error while generating token: ", err)
		return " ", err
	}
	return tokenString, nil
}
