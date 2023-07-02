package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

func main() {
	e0 := Endpoint{
		Path:        "/api/v1/ping",
		Description: "Ping/Pong health check",
		Method:      "GET",
	}
	e1 := Endpoint{
		Path:        "/api/v1/name",
		Description: "Set your name",
		Method:      "POST",
		Data:        struct{ Name string }{Name: "string"},
	}

	docs := Documentation{
		Endpoints: []Endpoint{e0, e1},
	}
	err := docs.Render()
	if err != nil {
		panic(err)
	}
	fmt.Println(docs.text.String())
}

const (
	Connect = "CONNECT"
	Delete  = "DELETE"
	Get     = "GET"
	Head    = "HEAD"
	Options = "OPTIONS"
	Patch   = "PATCH"
	Post    = "POST"
	Put     = "PUT"
	Trace   = "TRACE"
)

//Endpoint A descriptor for a REST endpoint
type Endpoint struct {
	Path        string
	Description string
	Method      string
	Data        interface{}
	Request     map[string]string
	Params      map[string]string
	Query       map[string]string
}

//ValidateMethod Validates the method is a REST Method
func (e *Endpoint) ValidateMethod() bool {
	m := strings.ToUpper(e.Method)
	switch m {
	case Connect, Delete, Get, Head, Options, Patch, Post, Put, Trace:
		return true
	default:
		return false
	}
}

//Documentation Structured data for compose the API Documentation
type Documentation struct {
	Endpoints []Endpoint
	text      *bytes.Buffer
}

//Render the html page
func (d *Documentation) Render() error {
	templater, err := template.ParseFiles("./static/index.html")
	if err != nil {
		return err
	}
	buff := new(bytes.Buffer)
	err = templater.Execute(buff, &struct{ Name string }{Name: "hi"})
	if err != nil {
		return err
	}
	d.text = buff
	return nil
}
