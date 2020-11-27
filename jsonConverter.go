package NestedJsonToStructConverter

import (
	"encoding/json"
	"log"
)

func StructuredJSON(nestedJsonData string, field string) string {

	data := dataFromNestedJson(nestedJsonData) // get mapped data

	fieldData := findFieldData(data, field) // find data for specified field

	byteData, err := json.Marshal(fieldData) // marshal it and return as string

	if err != nil {
		log.Fatal(err)
	}

	return string(byteData)
}

func dataFromNestedJson(data string) map[string]interface{} {

	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result) //marshal nestedJson to map.

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func findFieldData(data interface{}, name string) interface{} {

	// find if specified key is present in map or not in data.
	result, found := data.(map[string]interface{})[name]

	if !found {
		for key, value := range data.(map[string]interface{}) {

			/* if not found then loop through map data and check if given key present in nested loop or not.
			   check each key's type and get data according to that
			*/
			switch value.(type) {

			case []interface{}:
				// if it is of type []interface, again call same method for deep nesting.
				for _, j := range value.([]interface{}) {
					switch j.(type) {
					case map[string]interface{}, []interface{}:
						return findFieldData(j, name)
					}
				}

			case map[string]interface{}:
				/* if it is of type map[string]interface{}, then check if given key present.
				   if not present again call same method for find specified key. */
				result, found = value.(map[string]interface{})[name]
				if !found {
					return findFieldData(value, name)
				}

			case string, int, bool, float64:
				// if it is of type string, int, bool etc.., then return if specified key matched.
				if key == name {
					return value
				} else {
					continue
				}
			}
		}
	}

	return result
}
