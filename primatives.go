package main

import "strings"

var primatives = [24]string{
	//Numbers
	"number",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
	//Boolean
	"bool",
	//Strings
	"string",
	"uuid",
	"uri",
	"email",
	"byte", //base64
	//Dates
	"date",
	"datetime",
	//Complex
	"array",
	"object",
	"hashmap",
}

//DataType A given datatype that may be declared
type DataType struct {
	Name string
}

//Validate that the string is a valid datatype
func (d *DataType) Validate() bool {
	t := strings.ToLower(d.Name)
	valid := false
	for _, primative := range primatives {
		if t == primative {
			valid = true
			break
		}
	}

	return valid
}
