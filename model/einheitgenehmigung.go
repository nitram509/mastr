// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// EinheitenGenehmigung
type EinheitenGenehmigung struct {
	XMLName              xml.Name             `xml:"EinheitenGenehmigung"`
	EinheitenGenehmigung []EinheitGenehmigung `xml:"EinheitGenehmigung"`
}

type EinheitGenehmigung struct {
	XMLName                          xml.Name `xml:"EinheitGenehmigung"`
	GenMastrNummer                   string   `xml:"GenMastrNummer"`
	DatumLetzteAktualisierung        string   `xml:"DatumLetzteAktualisierung"`
	Art                              uint     `xml:"Art"` // referenziert Katalogwert
	Datum                            string   `xml:"Datum"`
	Behoerde                         string   `xml:"Behoerde"`
	Aktenzeichen                     string   `xml:"Aktenzeichen"`
	Frist                            string   `xml:"Frist"`
	Frist_nv                         string   `xml:"Frist_nv"`
	WasserrechtsNummer               string   `xml:"WasserrechtsNummer"`
	WasserrechtAblaufdatum           string   `xml:"WasserrechtAblaufdatum"`
	WasserrechtAblaufdatum_nv        string   `xml:"WasserrechtAblaufdatum_nv"`
	Registrierungsdatum              string   `xml:"Registrierungsdatum"`
	VerknuepfteEinheitenMaStRNummern string   `xml:"VerknuepfteEinheitenMaStRNummern"`
}