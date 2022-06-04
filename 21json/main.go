package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this will hide the password in the api
	Tags     []string `json:"tags,omitempty"` // if any field is nil, it will not show that
}

func main() {
	fmt.Println("Welcome to json in go")

	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 0, "lco.dev", "123abc", []string{"front-end", "web-dev"}},
		{"Python", 0, "lco.dev", "56df", []string{"back-end", "oop"}},
		{"Go", 100, "lco.sky", "jgj56", []string{"devops"}},
		{"c++", 100, "lco.sky", "jhuj56", nil},
	}

	// we have to package this as json data

	// finalJson, err := json.Marshal(lcoCourses) // use marshal indent to display in better way
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
