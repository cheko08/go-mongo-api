package main

//Address struct
type Address struct {
	StreetAddress string `json:streetAddress,omitempty`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	PostalCode    string `json:"postalCode,omitempty:`
}
