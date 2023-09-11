package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
	{
		"data" : {
			"object":"card",
			"id":"card_4734564656",
			"first_name":"Yasin",
			"last_name":"Kaplan",
			"balance":"54.950"
		}

	}
	`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)

	fmt.Println("----------")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
