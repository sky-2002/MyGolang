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

	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 0,
		"website": "lco.dev",
		"tags": ["front-end","web-dev"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was NOT VALID")
	}

	// adding data using key-value
	var myOnlineData map[string]interface{} // map keys are strings but values can be anything
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type of value is %T\n", k, v, v)
	}
}
