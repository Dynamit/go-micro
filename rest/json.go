package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//
// ParseJSON returns a struct from the encoded JSON received with an input request
func ParseJSON(r *http.Request, i interface{}) error {

	//
	// Get the JSON data from the input request
	defer r.Body.Close()
	var body, _ = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // limits to 1MB of input

	//
	// Decode the JSON data into our data format
	dec := json.NewDecoder(strings.NewReader(string(body)))
	for {
		if err := dec.Decode(&i); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	//
	// Return no errors
	return nil

}

//
// WriteJSON sends the JSON response to the responsewriter
func WriteJSON(w http.ResponseWriter, c int, i interface{}) error {

	//
	// Send the content type
	w.Header().Set("Content-Type", "application/json")

	//
	// Allow AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//
	// Send the HTTP response code
	w.WriteHeader(c)

	//
	// Encode and send the interface type
	if err := json.NewEncoder(w).Encode(i); err != nil {
		return err
	}

	//
	// Return no errors
	return nil

}
