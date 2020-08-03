package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func ValidateEmail(email string) (dot, arroba bool) {
	valid := 0
	for _, ok := range []byte(email) {
		if ok == 0x40 {
			arroba = true
			valid++
		}
		if ok == 0x2e {
			dot = true
			valid++
		}
	}
	if valid != 2 {
		arroba = false
		dot = false
	}

	return
}

func ReadJson() []User {

	jsonFile, err := os.Open("file.json")

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Users)
	return Users
}

func SaveUser(newJson User) (err error) {

	js := ReadJson()

	for _, v := range js {
		if v.Email == newJson.Email {
			err = errors.New("Existing email")
			return
		}
		/* block ip
		if v.IP == newJson.IP {
			err = errors.New("Existing IP LOL")
			return
		}*/
	}
	Users = append(Users, newJson)
	file, _ := json.MarshalIndent(Users, "", " ")
	err1 := ioutil.WriteFile("file.json", file, 0644)
	if err1 != nil {
		panic(err)
	}
	return nil
}
