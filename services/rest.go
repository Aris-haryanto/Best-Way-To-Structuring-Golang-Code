package services

// here's a file to handle request from http (rest API)
// this is of routing request from route and call service you need
// the routes, we define in main.go, so he in here we dont care about the framework we use
// you can use gin, mux, echo or something else

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RestServer struct {
	srvHello *Hello
}

// ====================

func (rest *RestServer) RegisterHello(acd *Hello) {
	rest.srvHello = acd
}

// ====================

func (rest *RestServer) SayHello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// read body request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	//set header
	w.Header().Add("Content-Type", "application/json")

	// convert request body to struct
	request := requestHello{}
	if err := json.Unmarshal(body, &request); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseHello{Code: 500, Message: err.Error(), Data: []byte{}})
		return
	}

	// set response
	result, _ := rest.srvHello.SaidHello(&request)
	w.WriteHeader(int(result.Code))
	json.NewEncoder(w).Encode(result)
}
