// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// Netze
type Netze struct {
	XMLName xml.Name `xml:"Netze"`
	Netze   []Netz   `xml:"Netz"`
}

type Netz struct {
	XMLName                   xml.Name `xml:"Netz"`
	DatumLetzteAktualisierung string   `xml:"DatumLetzteAktualisierung"`
	MastrNummer               string   `xml:"MastrNummer"`
	Sparte                    uint     `xml:"Sparte"` // referenziert Katalogwert
	KundenAngeschlossen       string   `xml:"KundenAngeschlossen"`
	GeschlossenesVerteilnetz  string   `xml:"GeschlossenesVerteilnetz"`
	Bezeichnung               string   `xml:"Bezeichnung"`
	Marktgebiet               uint     `xml:"Marktgebiet"` // referenziert Katalogwert

}