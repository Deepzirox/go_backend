package models

type User struct {
	Email string `json:"email,omitempty"`
	IP    string `json:email,omitempty`
}

type Candidates struct {
	Yami     int
	Bananero int
	Miguel   int
}

var Users []User
