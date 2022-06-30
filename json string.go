package main

import (
	"fmt"	
	"encoding/json"	
	)

var JSON = `{
	"name":"shikha",
	"jobtitle":"cyber security",
	"phone":{
		"home":"123-466-789",
		"office":"564"
	},
	"email":"shikha@gmail.com"
}`

func main() {
	
	var info map[string]interface{}
	json.Unmarshal([]byte(JSON),&info)

	fmt.Println(info["name"])
	fmt.Println(info["jobtitle"])
	fmt.Println(info["email"])
	fmt.Println(info["phone"].(map[string]interface{})["home"])	
	fmt.Println(info["phone"].(map[string]interface{})["office"])
}