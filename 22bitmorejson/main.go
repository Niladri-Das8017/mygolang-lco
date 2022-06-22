package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //no space allooweedddd while creating json alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //"-" indecates that, we don't want it to be reflected
	Tags     []string `json:"tags,omitempty"` //it will ommit the field if the data is empty
}

func main() {
	fmt.Println("\nWelcome to json video!!!")
	//EncodeJson()
	DecodeeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "Youtube", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "Youtube", "bcd123", []string{"full-stack", "js"}},
		{"Angular", 99, "Youtube", "efg123", nil},
	}

	//Package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeeJson() {
	jesonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS",
			"Price": 299,
			"website": "Youtube",
			"tags": ["web-dev","js"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jesonDataFromWeb)

	if checkValid {

		fmt.Println("JSON was valid")
		json.Unmarshal(jesonDataFromWeb, &lcoCourse) //we don't wan to pass a copy, hence the reference
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	
	// some cases where you just want a key value pair
	var myOnlineData map[string]interface{} //it can be array, int or anything, hence we're using interface
	json.Unmarshal(jesonDataFromWeb, &myOnlineData) //we don't wan to pass a copy, hence the reference
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}

}
