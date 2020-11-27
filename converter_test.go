package NestedJsonToStructConverter

import (
	"testing"
)

var testCases = []struct {
	name     string
	jsonData string
	key      string
	expected interface{}
}{
	{"simpleJSONData",
		"{\"data\" : {\"student\" : {\"name\" : \"abc\",\"std\" : \"8\",\"roll no\" : \"20\"}}}",
		"student",
		"\"student\""},

	{"nestedStudentNullData",
		"{\"data\" : {\"student\" : {\"name\" : \"abc\",\"std\" : \"8\",\"roll no\" : \"20\"," +
			"\"subjects\" : [{\"name\" : \"english\",\"tutor\" : \"xyz sir\"},{\"name\" : \"maths\",\"tutor\" : \"pqr madam\"}]}}}",
		"nam",
		"null"},

	{"nestedStudentData",
		"{\"data\" : {\"student\" : {\"name\" : \"abc\",\"std\" : \"8\",\"roll no\" : \"20\"," +
			"\"subjects\" : [{\"name\" : \"english\",\"tutor\" : \"xyz sir\"},{\"name\" : \"maths\",\"tutor\" : \"pqr madam\"}]}}}",
		"name",
		"\"abc\""},

	{"nestedRecipeData",
		"{\"items\":{\"item\":[{\"id\": \"0001\",\"type\": \"donut\",\"name\": \"Cake\",\"ppu\": 0.55," +
			"\"batters\":{\"batter\":[{ \"id\": \"1001\", \"type\": \"Regular\" },{ \"id\": \"1002\", \"type\": \"Chocolate\" },{ \"id\": \"1003\", \"type\": \"Blueberry\" },{ \"id\": \"1004\", \"type\": \"Devil's Food\" }]}," +
			"\"topping\":[{ \"id\": \"5001\", \"type\": \"None\" },{ \"id\": \"5002\", \"type\": \"Glazed\" }," +
			"{ \"id\": \"5005\", \"type\": \"Sugar\" },{ \"id\": \"5007\", \"type\": \"Powdered Sugar\" },{ \"id\": \"5006\", \"type\": \"Chocolate with Sprinkles\" }," +
			"{ \"id\": \"5003\", \"type\": \"Chocolate\" },{ \"id\": \"5004\", \"type\": \"Maple\" }]}]}}",
		"batters",
		"{\"batter\":[{\"id\":\"1001\",\"type\":\"Regular\"},{\"id\":\"1002\",\"type\":\"Chocolate\"},{\"id\":\"1003\",\"type\":\"Blueberry\"},{\"id\":\"1004\",\"type\":\"Devil's Food\"}]}"},
}

func TestJSONConverter(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if output := StructuredJSON(test.jsonData, test.key); test.expected != output {
				t.Errorf("For input: %v, expected: %v, but got: %v", test.key, test.expected, output)
			}
		})
	}
}
