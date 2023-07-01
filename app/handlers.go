package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json: "full_name" xml: "full_name"` //if we dont use these tags then response will be like [{"Name":"Ashish","City":"Ranchi","ZipCode":"834009"},{"Name":"Ramesh","City":"Bangalore","ZipCode":"560017"}]
	City    string `json: "city" xml: "city"`           //which is in capital letters
	ZipCode string `json: "zip_code" xml: "zip_code"`
}

// in handler func we have two params, response writer that writes back to client and second is  request reader that takes request from client
// we will use Fprint func to send the response to the client, this takes two params: one is IO writer and other is data that we want to send to IO writer
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!") //simple handler func that returns "Hello World!" as response in http ResponseWriter
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) { // Handler func to list all customers
	customer := []Customer{
		{"Ashish", "Ranchi", "834009"},
		{"Ramesh", "Bangalore", "560017"},
	}

	if r.Header.Get("Content-Type") == "application/xml" { //if request Header contains Content-Type as xml then
		w.Header().Add("Content-Type", "application/xml") // we want to give output as xml
		xml.NewEncoder(w).Encode(customer)
	} else { //else we will give response as json
		w.Header().Add("Content-Type", "application/json") // this is to recognize the response content type as json,
		//without this the response will be of type json only since we have used json encoder below but wouldn't be viewed as json,
		//instead will be viewed as raw text in postman response

		json.NewEncoder(w).Encode(customer) //encode the list to json format
	}
}
