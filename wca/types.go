package main

type Member struct {
	Name           string
	Branches       map[string][]string
	EnrolledSince  string
	Profile        string
	Address        string
	ContactDetails map[string]string
	OfficeContacts []map[string]string
}

type CountryCodes struct {
	Codes []string `json:"country_codes"`
}
