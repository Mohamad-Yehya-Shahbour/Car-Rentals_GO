package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
	"time"
	jwt "github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ResponseHandler(w http.ResponseWriter, msgKey string, msg string, status int){
	resp := make(map[string]string)
		resp[msgKey] = msg
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(status)
		w.Write(jsonResp)
}


func GenerateJWT() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = "mohamad yehya"
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Printf("Something Went Wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("there was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                log.Fatal(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {
			w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, "Not Authorized")
        }
    })
}