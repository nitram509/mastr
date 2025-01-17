// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// Marktrollen
type Marktrollen struct {
	XMLName     xml.Name     `xml:"Marktrollen"`
	Marktrollen []Marktrolle `xml:"Marktrolle"`
}

type Marktrolle struct {
	XMLName                              xml.Name `xml:"Marktrolle"`
	MarktakteurMastrNummer               string   `xml:"MarktakteurMastrNummer"` // referenziert Marktakteur
	DatumLetzteAktualisierung            string   `xml:"DatumLetzteAktualisierung"`
	MastrNummer                          string   `xml:"MastrNummer"`
	Marktrolle                           string   `xml:"Marktrolle"`
	BundesnetzagenturBetriebsnummer      string   `xml:"BundesnetzagenturBetriebsnummer"`
	BundesnetzagenturBetriebsnummer_nv   string   `xml:"BundesnetzagenturBetriebsnummer_nv"`
	Marktpartneridentifikationsnummer    string   `xml:"Marktpartneridentifikationsnummer"`
	Marktpartneridentifikationsnummer_nv string   `xml:"Marktpartneridentifikationsnummer_nv"`
	KontaktdatenMarktrolle               string   `xml:"KontaktdatenMarktrolle"`
}
