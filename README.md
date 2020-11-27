## JSONToStructConverter

##### This repository gives you data of your choice from complex nested JSON in golang.
##### Call ****StructuredJSON(JSONData, fieldName)**** and see the result.

### Example

````
jsonData := "{\"data\" : {\"student\" : {\"name\" : \"abc\",\"std\" : \"8\",\"roll no\" : \"20\"}}}"

StructuredJSON(jsonData, "data")

Result : {"name":"abc","roll no":"20","std":"8"}
````

##### We can unmarshal result to required struct.