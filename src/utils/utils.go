package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
)

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
