// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// Einheitentypen
type Einheitentypen struct {
	XMLName        xml.Name       `xml:"Einheitentypen"`
	Einheitentypen []Einheitentyp `xml:"Einheitentyp"`
}

type Einheitentyp struct {
	XMLName xml.Name `xml:"Einheitentyp"`
	Id      uint     `xml:"Id"`
	Wert    string   `xml:"Wert"`
}