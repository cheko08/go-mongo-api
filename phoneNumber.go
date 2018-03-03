package main

//PhoneNumber struct
type PhoneNumber struct {
	Number      string `json:phoneNumber,omitempty`
	CountryCode string `json:countryCode,omitempty`
}
