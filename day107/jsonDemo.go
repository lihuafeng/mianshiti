package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//bool, for JSON booleans
	//float64, for JSON numbers
	//string, for JSON strings
	//[]interface{}, for JSON arrays
	//map[string]interface{}, for JSON objects
	//nil for JSON null
	var p *int
	fmt.Println(p)
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
	fmt.Println(string(b))
}
