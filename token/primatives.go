package token

import "strings"

type Primative string

var primatives = [24]Primative{
	//Numbers
	"number",
	"int",
	"i8",
	"i16",
	"i32",
	"i64",
	"uint",
	"u8",
	"u16",
	"u32",
	"u64",
	"f32",
	"f64",
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
type Data struct {
	Type Primative
	Name string
}

//Validate that the string is a valid datatype
func (d *Data) Validate() bool {
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
